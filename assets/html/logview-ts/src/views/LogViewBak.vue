<template>
  <div class="homewrap">
    <el-container class="home-container">
      <el-aside width="20%" v-loading="loading">

        <div
            style="text-align: left;border: solid 1px #d9dede;;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);padding: 5px;margin-top: 1px;">
          <el-form>
            <el-form-item label="ws地址：">
              <!--              <el-input-->
              <!--                  size="mini"-->
              <!--                  v-model="datas.localaddr"-->
              <!--                  placeholder="请输入ws地址"-->
              <!--              ></el-input>-->
              <el-select
                  :disabled="connStatus"
                  v-model="datas.localaddr"
                  placeholder="请输入ws地址"
                  clearable
                  filterable
                  @blur="selectBlur"
                  @clear="selectClear"
                  @change="selectChange"
              >
                <el-option
                    v-for="(item,index) in options"
                    :key="index"
                    :label="item.label"
                    :value="item.value"/>
              </el-select>
            </el-form-item>

            <el-form-item label="日志路径：">
              <el-input
                  size="mini"
                  v-model="datas.logpath"
                  placeholder="请输日志路径"
              ></el-input>

            </el-form-item>


            <div style="display: flex;width: 100%;margin-top: 5px">
              <el-button :type="datas.btncolor" style="font-size: 10px; width: 100%;" @click="onStart()">
                {{ datas.btntext }}
              </el-button>
              <el-button size="mini" style="font-size: 10px;width: 100%;" @click="onRefreshLog()">刷新日志</el-button>
              <el-button size="mini" style="font-size: 10px;width: 100%;" @click="onClearLog()">清空日志</el-button>
            </div>

          </el-form>
        </div>


        <!--文件结构-->
        <div
            style="text-align: left;border: solid 1px #d9dede;;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);padding: 5px;margin-top: 15px;">
          <div class="border-title">
            <span>文件结构</span>
          </div>

          <el-tree-v2 :data="treeData" :props="props" :height="500"/>
        </div>


      </el-aside>
      <el-container>
        <div style="width: 100%">
          <div
              style="
              border: solid 1px #d9dede;
              box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12),
                0 0 6px rgba(0, 0, 0, 0.04);
              padding: 5px;
              width: 100%;
              height: 90%;
              overflow: auto;
              word-break: break-all;
            "
          >
            <!-- 日志显示区域 -->
            <div
                ref="txtContent"
                id="txtContent"
                style="
                text-align: left;
                border: 1px solid #646060;
                height: 100%;
                overflow: auto;
                word-break: break-all;
              "
            />
          </div>
        </div>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts">
import {Options, Vue} from 'vue-class-component';
import {ElMessage, ElMessageBox} from 'element-plus'
import type {Action} from 'element-plus'

interface Tree {
  id: string
  label: string
  children?: Tree[]
}

const getKey = (prefix: string, id: number) => {
  return `${prefix}-${id}`
}

const createData = (maxDeep: number, maxChildren: number, minNodesNumber: number, deep = 1, key = 'node'): Tree[] => {
  let id = 0
  return Array.from({length: minNodesNumber})
      .fill(deep)
      .map(() => {
        const childrenNumber =
            deep === maxDeep ? 0 : Math.round(Math.random() * maxChildren)
        const nodeKey = getKey(key, ++id)
        return {
          id: nodeKey,
          label: nodeKey,
          children: childrenNumber
              ? createData(maxDeep, maxChildren, childrenNumber, deep + 1, nodeKey)
              : undefined,
        }
      })
}


@Options({
  props: {
    msg: String
  }
})
export default class LogView extends Vue {
  msg!: string
  websock: any;

  private props = {
    value: 'id',
    label: 'label',
    children: 'children',
  }

  private loading = false;
  private connStatus = false;
  private options = [{
    value: 'wss://cyjy-iot.chengyang.gov.cn/clink/gtbx/logws',
    label: '城阳环境'
  }, {value: 'ws://iot.clife.net:31307/echo', label: 'clife生产环境'}]
  private treeData = createData(2, 2, 3)
  private datas = {
    btntext: "打开",
    btncolor: "primary",
    logpath: "/logs",
    localaddr: "",//"ws://" + window.location.host
    websock: null,
    logDiv: document.getElementById("txtContent"),
  };
  private wshost = '';
  private apihost = window.location.protocol;

  public selectBlur(e: any) {
    // 意见类型
    if (e.target.value !== '') {
      this.datas.localaddr = e.target.value;
      this.showLog(this.datas.localaddr)
      this.$forceUpdate();   // 强制更新
    }
  }


  public selectClear() {
    this.datas.localaddr = '';
    this.$forceUpdate();
  }

  public selectChange(val: any): void {
    this.datas.localaddr = val;
    this.$forceUpdate();
    this.showLog(this.datas.localaddr)
  }

  public onStart(): void {
    if (this.datas.localaddr === '') {
      ElMessageBox.alert('请填写ws地址哦～～', '警告', {
        // if you want to disable its autofocus
        // autofocus: false,
        confirmButtonText: 'OK',
        callback: (action: Action) => {
          ElMessage({
            type: 'info',
            message: `action: ${action}`,
          })
        },
      })
      return console.log("地址是空的哦")
    }
    console.log("打开")
    this.loading = true;
    if (!this.connStatus) {
      this.initWebSocket()
    } else {
      this.websock.close()
    }
  }

  public initWebSocket(): void {
    console.log("initWebSocket")
    if (typeof WebSocket === 'undefined')
      return console.log('您的浏览器不支持websocket')
    // let host = `ws://${this.localaddr}/echo`
    //ws://api.uuxia.cn/gows
    // let host = `ws://localhost:8080/echo`
    // let host = `ws://api.uuxia.cn/gows`
    console.log("host", this.datas.localaddr);
    try {
      this.websock = new WebSocket(this.datas.localaddr)
      this.websock.onopen = (e: any) => {
        this.datas.btncolor = "danger";
        this.loading = false;
        this.connStatus = true
        console.log(e.currentTarget.url)
        console.log('websocket connect sucessully..', e)
        this.showLog('连接成功 ' + e.currentTarget.url)
        this.datas.btntext = "关闭"
      }
      this.websock.onmessage = (param: MessageEvent) => {
        console.log('onmessage', param)
        this.showLog(param.data)
      }

      this.websock.onclose = (e: any) => {
        this.datas.btncolor = "primary";
        this.loading = false;
        console.log('onclose received a message', e);
        this.showLog("连接关闭:" + JSON.stringify(e));
        this.connStatus = false
        this.datas.btntext = "打开"
      };

      this.websock.onerror = (e: any) => {
        this.datas.btncolor = "primary";
        this.loading = false;
        console.log('onerror received a message', e);
        this.showLog("连接错误:" + JSON.stringify(e));
        this.connStatus = false
        this.datas.btntext = "打开"
      };

    } catch (e) {
      console.log('catch received a message', e);
    }
  }

  public showLog(e: string | Node): void {
    var p = document.createElement('div')
    p.append(e)
    if (this.datas.logDiv != null) {
      this.datas.logDiv.append(p)
      this.datas.logDiv.scrollTop = this.datas.logDiv.scrollHeight;
      this.datas.logDiv.scrollIntoView();
    }
  }

  public onClearLog(): void {
    if (this.datas.logDiv != null) {
      this.datas.logDiv.innerText = "";
    }
  }

  public onRefreshLog(): void {
    this.fetchData(this.datas.logpath)
  }
  public get allname() {
    return 'computed ' + this.msg;
  }


  public fetchData(path : any): void {
    fetch(this.apihost + 'api/files?path=' + path, {
      credentials: 'include'
    })
        .then((res) => {
          return res.json()
        })
        .then((json) => {
          // status.value = new Array()
          // for (let key in json) {
          //   for (let ps of json[key]) {
          //     console.log(ps)
          //     status.value.push(ps)
          //   }
          // }

          this.treeData = json

          this.showLog(JSON.stringify(json))
        })
        .catch((err) => {
          ElMessage({
            showClose: true,
            message: 'Get status info from frpc failed!' + err,
            type: 'warning',
          })
        })
  }

  public created(): void {
    console.log('created');
  }

  public mounted(): void {
    console.log('mounted',window.location)
    console.log('host',window.location.host)
    console.log('origin',window.location.origin)
    console.log('pathname',window.location.pathname)
    console.log('protocol',window.location.protocol)


    let url = window.location.pathname;
    let list = url.split('/')
    console.log('list',list); // Banana, Ma
    let api = '/';
    list.forEach((value, index, array) => {
      if (index > 0 && index < array.length - 2){
        console.log('forEach',value,index,array.length); // Banana, Ma
        api = api.concat(value,"/")
      }
    });
    console.log('api',api); // Banana, Ma
    if (window.location.protocol.startsWith('http')){
      this.wshost = this.wshost.concat("ws://",window.location.hostname,api)
    }else{
      this.wshost = this.wshost.concat("wss://",window.location.hostname,api)
    }
    this.wshost = this.wshost.concat("echo")
    this.apihost = this.apihost.concat("//",window.location.hostname,api)

    console.log('wshost',this.wshost); // Banana, Ma
    console.log('apihost',this.apihost); // Banana, Ma


    this.datas.logDiv = document.getElementById("txtContent");
    this.showLog(JSON.stringify(this.treeData))
  }

}
</script>

<style scoped lang="scss">
.homewrap {
  position: absolute;
  top: 0;
  left: 0;
  width: 99%;
  height: 100%;
  text-align: center;
}

.el-container {
  height: 100%;
}

.el-container .el-form-item {
  margin-bottom: 10px;
}

.el-aside {
  background-color: #ffffff;
  padding: 10px;
}

.el-main {
  background-color: #a0dce6;
}

.el-tag {
  background-color: #409eff;
  width: 100%;
  color: #ffffff;
  font-size: 18px;
  margin-bottom: 4px;
  text-align: center;
}

.border-title {
  width: 90px;
  font-size: 13px;
  background: white;
  text-align: center;
  margin-top: -12px;
  margin-left: 10px;
}
</style>
