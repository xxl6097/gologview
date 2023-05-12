(function(){"use strict";var t={9602:function(t,o,e){var n=e(9242),s=e(3396);function i(t,o,e,n,i,l){const a=(0,s.up)("LogView");return(0,s.wg)(),(0,s.iD)("div",null,[(0,s.Wm)(a)])}var l=e(7139);const a=t=>((0,s.dD)("data-v-5ecdb2fe"),t=t(),(0,s.Cn)(),t),r={class:"homewrap"},c={style:{"text-align":"left",border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)",padding:"5px","margin-top":"1px"}},d={style:{display:"flex",width:"100%","margin-top":"5px"}},h={style:{"text-align":"left",border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)",padding:"5px","margin-top":"15px"}},u=a((()=>(0,s._)("div",{class:"border-title"},[(0,s._)("span",null,"文件结构")],-1))),g={style:{width:"100%"}},p={style:{border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, 0.12),\n                0 0 6px rgba(0, 0, 0, 0.04)",padding:"5px",width:"100%",height:"90%",overflow:"auto","word-break":"break-all"}},w={ref:"txtContent",id:"txtContent",style:{"text-align":"left",border:"1px solid #646060",height:"100%",overflow:"auto","word-break":"break-all"}},f={class:"rightMenu"};function m(t,o,e,i,a,m){const b=(0,s.up)("el-input"),v=(0,s.up)("el-form-item"),x=(0,s.up)("el-button"),y=(0,s.up)("el-form"),k=(0,s.up)("el-tree-v2"),_=(0,s.up)("el-aside"),W=(0,s.up)("el-container"),O=(0,s.up)("CirclePlus"),S=(0,s.up)("el-icon"),C=(0,s.up)("CircleClose"),D=(0,s.Q2)("loading");return(0,s.wg)(),(0,s.iD)("div",r,[(0,s.Wm)(W,{class:"home-container"},{default:(0,s.w5)((()=>[(0,s.wy)(((0,s.wg)(),(0,s.j4)(_,{width:"20%"},{default:(0,s.w5)((()=>[(0,s._)("div",c,[(0,s.Wm)(y,null,{default:(0,s.w5)((()=>[(0,s.Wm)(v,{label:"ws地址："},{default:(0,s.w5)((()=>[(0,s.Wm)(b,{size:"mini",modelValue:t.wshost,"onUpdate:modelValue":o[0]||(o[0]=o=>t.wshost=o),placeholder:"请输入ws地址"},null,8,["modelValue"])])),_:1}),(0,s.Wm)(v,{label:"日志路径："},{default:(0,s.w5)((()=>[(0,s.Wm)(b,{size:"mini",modelValue:t.datas.logpath,"onUpdate:modelValue":o[1]||(o[1]=o=>t.datas.logpath=o),placeholder:"请输日志路径"},null,8,["modelValue"])])),_:1}),(0,s._)("div",d,[(0,s.Wm)(x,{type:t.datas.btncolor,style:{"font-size":"10px",width:"100%"},onClick:o[2]||(o[2]=o=>t.onStart())},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.datas.btntext),1)])),_:1},8,["type"]),(0,s.Wm)(x,{size:"mini",style:{"font-size":"10px",width:"100%"},onClick:o[3]||(o[3]=o=>t.onRefreshLog())},{default:(0,s.w5)((()=>[(0,s.Uk)("刷新日志")])),_:1}),(0,s.Wm)(x,{size:"mini",style:{"font-size":"10px",width:"100%"},onClick:o[4]||(o[4]=o=>t.onClearLog())},{default:(0,s.w5)((()=>[(0,s.Uk)("清空日志")])),_:1})])])),_:1})]),(0,s._)("div",h,[u,(0,s.Wm)(k,{data:t.treeData,props:t.props,height:500},null,8,["data","props"])])])),_:1})),[[D,t.loading]]),(0,s.Wm)(W,null,{default:(0,s.w5)((()=>[(0,s._)("div",g,[(0,s._)("div",p,[(0,s._)("div",w,null,512)])])])),_:1})])),_:1}),(0,s.wy)((0,s._)("div",f,[(0,s._)("ul",null,[(0,s._)("li",{onClick:o[5]||(o[5]=(...o)=>t.addMenu&&t.addMenu(...o))},[(0,s.Wm)(S,null,{default:(0,s.w5)((()=>[(0,s.Wm)(O)])),_:1}),(0,s.Uk)("新建同级 ")]),(0,s._)("li",{onClick:o[6]||(o[6]=(...o)=>t.addSonMenu&&t.addSonMenu(...o))},[(0,s.Wm)(S,null,{default:(0,s.w5)((()=>[(0,s.Wm)(O)])),_:1}),(0,s.Uk)("新建子级 ")]),(0,s._)("li",{onClick:o[7]||(o[7]=(...o)=>t.delMenu&&t.delMenu(...o))},[(0,s.Wm)(S,null,{default:(0,s.w5)((()=>[(0,s.Wm)(C)])),_:1}),(0,s.Uk)("删除功能 ")])])],512),[[n.F8,t.showRightMenu]])])}var b=e(7327),v=e(6520),x=e(2655),y=e(7178),k=function(t,o,e,n){var s,i=arguments.length,l=i<3?o:null===n?n=Object.getOwnPropertyDescriptor(o,e):n;if("object"===typeof Reflect&&"function"===typeof Reflect.decorate)l=Reflect.decorate(t,o,e,n);else for(var a=t.length-1;a>=0;a--)(s=t[a])&&(l=(i<3?s(l):i>3?s(o,e,l):s(o,e))||l);return i>3&&l&&Object.defineProperty(o,e,l),l};const _=(t,o)=>`${t}-${o}`,W=(t,o,e,n=1,s="node")=>{let i=0;return Array.from({length:e}).fill(n).map((()=>{const e=n===t?0:Math.round(Math.random()*o),l=_(s,++i);return{id:l,label:l,children:e?W(t,o,e,n+1,l):void 0}}))};let O=class extends v.w3{constructor(...t){super(...t),(0,b.Z)(this,"msg",void 0),(0,b.Z)(this,"websock",void 0),(0,b.Z)(this,"props",{value:"id",label:"label",children:"children"}),(0,b.Z)(this,"loading",!1),(0,b.Z)(this,"connStatus",!1),(0,b.Z)(this,"options",[{value:"wss://cyjy-iot.chengyang.gov.cn/clink/gtbx/logws",label:"城阳环境"},{value:"ws://iot.clife.net:31307/echo",label:"clife生产环境"}]),(0,b.Z)(this,"treeData",W(2,2,3)),(0,b.Z)(this,"datas",{btntext:"打开",btncolor:"primary",logpath:"/",websock:null,logDiv:document.getElementById("txtContent")}),(0,b.Z)(this,"wshost",""),(0,b.Z)(this,"apihost",window.location.origin)}onStart(){if(""===this.wshost)return x.T.alert("请填写ws地址哦～～","警告",{confirmButtonText:"OK",callback:t=>{(0,y.z8)({type:"info",message:`action: ${t}`})}}),console.log("地址是空的哦");console.log("打开"),this.loading=!0,this.connStatus?this.websock.close():this.initWebSocket()}initWebSocket(){if(console.log("initWebSocket"),"undefined"===typeof WebSocket)return console.log("您的浏览器不支持websocket");console.log("host",this.wshost);try{this.websock=new WebSocket(this.wshost),this.websock.onopen=t=>{this.datas.btncolor="danger",this.loading=!1,this.connStatus=!0,console.log(t.currentTarget.url),console.log("websocket connect sucessully..",t),this.showLog("连接成功 "+t.currentTarget.url),this.datas.btntext="关闭"},this.websock.onmessage=t=>{console.log("onmessage",t),this.showLog(t.data)},this.websock.onclose=t=>{this.datas.btncolor="primary",this.loading=!1,console.log("onclose received a message",t),this.showLog("连接关闭:"+JSON.stringify(t)),this.connStatus=!1,this.datas.btntext="打开"},this.websock.onerror=t=>{this.datas.btncolor="primary",this.loading=!1,console.log("onerror received a message",t),this.showLog("连接错误:"+JSON.stringify(t)),this.connStatus=!1,this.datas.btntext="打开"}}catch(t){console.log("catch received a message",t)}}showLog(t){var o=document.createElement("div");o.append(t),null!=this.datas.logDiv&&(this.datas.logDiv.append(o),this.datas.logDiv.scrollTop=this.datas.logDiv.scrollHeight,this.datas.logDiv.scrollIntoView())}onClearLog(){null!=this.datas.logDiv&&(this.datas.logDiv.innerText="")}onRefreshLog(){this.fetchData(this.datas.logpath)}get allname(){return"computed "+this.msg}fetchData(t){fetch(this.apihost+"api/files?path="+t,{credentials:"include"}).then((t=>(this.showLog(JSON.stringify(t)),t.json()))).then((t=>{this.treeData=t,this.showLog(JSON.stringify(t))})).catch((t=>{(0,y.z8)({showClose:!0,message:"Get status failed!"+t,type:"warning"}),this.showLog(t)}))}created(){console.log("created")}mounted(){console.log("mounted",window.location),console.log("host",window.location.host),console.log("origin",window.location.origin),console.log("pathname",window.location.pathname),console.log("protocol",window.location.protocol);let t=window.location.pathname,o=t.split("/");console.log("list",o);let e="/";o.forEach(((t,o,n)=>{o>0&&o<n.length-2&&(console.log("forEach",t,o,n.length),e=e.concat(t,"/"))})),console.log("api",e),window.location.protocol.startsWith("http")?this.wshost=this.wshost.concat("ws://",window.location.host,e):this.wshost=this.wshost.concat("wss://",window.location.host,e),this.wshost=this.wshost.concat("echo"),console.log("wshost",this.wshost),this.apihost=this.apihost.concat(e),console.log("apihost",this.apihost),this.datas.logDiv=document.getElementById("txtContent"),this.showLog(JSON.stringify(this.treeData))}};O=k([(0,v.Ei)({props:{msg:String}})],O);var S=O,C=e(89);const D=(0,C.Z)(S,[["render",m],["__scopeId","data-v-5ecdb2fe"]]);var L=D,Z=(0,s.aZ)({name:"HomeView",components:{LogView:L}});const j=(0,C.Z)(Z,[["render",i]]);var M=j,z=e(2777);e(3942),e(4415);(0,n.ri)(M).use(z.Z).mount("#app")}},o={};function e(n){var s=o[n];if(void 0!==s)return s.exports;var i=o[n]={exports:{}};return t[n].call(i.exports,i,i.exports,e),i.exports}e.m=t,function(){var t=[];e.O=function(o,n,s,i){if(!n){var l=1/0;for(d=0;d<t.length;d++){n=t[d][0],s=t[d][1],i=t[d][2];for(var a=!0,r=0;r<n.length;r++)(!1&i||l>=i)&&Object.keys(e.O).every((function(t){return e.O[t](n[r])}))?n.splice(r--,1):(a=!1,i<l&&(l=i));if(a){t.splice(d--,1);var c=s();void 0!==c&&(o=c)}}return o}i=i||0;for(var d=t.length;d>0&&t[d-1][2]>i;d--)t[d]=t[d-1];t[d]=[n,s,i]}}(),function(){e.n=function(t){var o=t&&t.__esModule?function(){return t["default"]}:function(){return t};return e.d(o,{a:o}),o}}(),function(){e.d=function(t,o){for(var n in o)e.o(o,n)&&!e.o(t,n)&&Object.defineProperty(t,n,{enumerable:!0,get:o[n]})}}(),function(){e.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){e.o=function(t,o){return Object.prototype.hasOwnProperty.call(t,o)}}(),function(){var t={143:0};e.O.j=function(o){return 0===t[o]};var o=function(o,n){var s,i,l=n[0],a=n[1],r=n[2],c=0;if(l.some((function(o){return 0!==t[o]}))){for(s in a)e.o(a,s)&&(e.m[s]=a[s]);if(r)var d=r(e)}for(o&&o(n);c<l.length;c++)i=l[c],e.o(t,i)&&t[i]&&t[i][0](),t[i]=0;return e.O(d)},n=self["webpackChunkts_logview"]=self["webpackChunkts_logview"]||[];n.forEach(o.bind(null,0)),n.push=o.bind(null,n.push.bind(n))}();var n=e.O(void 0,[998],(function(){return e(9602)}));n=e.O(n)})();