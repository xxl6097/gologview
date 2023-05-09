(function(){"use strict";var e={9804:function(e,o,t){var n=t(9242),i=t(3396);function a(e,o,t,n,a,r){const s=(0,i.up)("LogView");return(0,i.wg)(),(0,i.iD)("div",null,[(0,i._)("div",null,[(0,i.Wm)(s)])])}var r=t(7139);const s={class:"homewrap"},l={style:{display:"flex"}},c={style:{width:"100%"}},u={style:{border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)",padding:"5px",width:"100%",height:"90%",overflow:"auto","word-break":"break-all"}},d={ref:"txtContent",id:"txtContent",style:{"text-align":"left",border:"1px solid #646060",height:"100%",overflow:"auto","word-break":"break-all"}};function f(e,o,t,n,a,f){const h=(0,i.up)("el-tag"),g=(0,i.up)("el-aside"),v=(0,i.up)("el-container");return(0,i.wg)(),(0,i.iD)("div",s,[(0,i.Wm)(v,{class:"home-container"},{default:(0,i.w5)((()=>[(0,i.Wm)(g,{width:"10%"},{default:(0,i.w5)((()=>[(0,i._)("div",l,[(0,i.Wm)(h,{style:{"font-size":"15px","font-weight":"bold"},size:"large",effect:"dark"},{default:(0,i.w5)((()=>[(0,i.Uk)((0,r.zw)(e.names.title),1)])),_:1})])])),_:1}),(0,i.Wm)(v,null,{default:(0,i.w5)((()=>[(0,i._)("div",c,[(0,i._)("div",u,[(0,i._)("div",d,null,512)])])])),_:1})])),_:1})])}var h=t(3671);let g={name:"LogView",data(){return{websock:null,data:{localport:18899},names:{title:"TCP/IP调试工具"},sData1:"",div:{logDiv:document.getElementById("txtContent")}}},computed:{},mounted(){this.div.logDiv=document.getElementById("txtContent"),this.initWebSocket()},methods:{left:h.left,onWorkModeChange(e){console.log("aaaaa",e),this.showLog(e)},onRecvHexDisplay(e){console.log("aaaaa",e),this.showLog(e)},onSendHexDisplay(e){console.log("aaaaa",e),this.showLog(e)},onIssueChange(e){console.log("aaaaa",e),this.showLog(e)},onIntervarChange(e){console.log("aaaaa",e),this.showLog(e)},onTcpSelectChange(e){console.log("aaaaa",e),this.showLog(e)},onStart(e){console.log("aaaaa",e),this.showLog(e);let o={message:"sys_info"};this.websocketsend(JSON.stringify(o))},showLog(e){var o=document.createElement("div");o.append(e),this.div.logDiv.append(o),this.div.logDiv.scrollTop=this.div.logDiv.scrollHeight,this.div.logDiv.scrollIntoView()},websocketsend(e){this.websock.send(e)},initWebSocket(){if("undefined"===typeof WebSocket)return console.log("您的浏览器不支持websocket");let e=this;this.websock=new WebSocket("ws://localhost:8081/echo"),this.websock.onmessage=function(o){console.log("onmessage",o),e.showLog(o.data)},this.websock.onopen=function(e){console.log("onopen",e)}}}};var v=g,p=t(89);const w=(0,p.Z)(v,[["render",f]]);var b=w,m={name:"App",components:{LogView:b},methods:{}};const k=(0,p.Z)(m,[["render",a]]);var x=k,y=t(9906);t(3942),t(4415);(0,n.ri)(x).use(y.Z).mount("#app")}},o={};function t(n){var i=o[n];if(void 0!==i)return i.exports;var a=o[n]={exports:{}};return e[n].call(a.exports,a,a.exports,t),a.exports}t.m=e,function(){var e=[];t.O=function(o,n,i,a){if(!n){var r=1/0;for(u=0;u<e.length;u++){n=e[u][0],i=e[u][1],a=e[u][2];for(var s=!0,l=0;l<n.length;l++)(!1&a||r>=a)&&Object.keys(t.O).every((function(e){return t.O[e](n[l])}))?n.splice(l--,1):(s=!1,a<r&&(r=a));if(s){e.splice(u--,1);var c=i();void 0!==c&&(o=c)}}return o}a=a||0;for(var u=e.length;u>0&&e[u-1][2]>a;u--)e[u]=e[u-1];e[u]=[n,i,a]}}(),function(){t.n=function(e){var o=e&&e.__esModule?function(){return e["default"]}:function(){return e};return t.d(o,{a:o}),o}}(),function(){t.d=function(e,o){for(var n in o)t.o(o,n)&&!t.o(e,n)&&Object.defineProperty(e,n,{enumerable:!0,get:o[n]})}}(),function(){t.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){t.o=function(e,o){return Object.prototype.hasOwnProperty.call(e,o)}}(),function(){var e={143:0};t.O.j=function(o){return 0===e[o]};var o=function(o,n){var i,a,r=n[0],s=n[1],l=n[2],c=0;if(r.some((function(o){return 0!==e[o]}))){for(i in s)t.o(s,i)&&(t.m[i]=s[i]);if(l)var u=l(t)}for(o&&o(n);c<r.length;c++)a=r[c],t.o(e,a)&&e[a]&&e[a][0](),e[a]=0;return t.O(u)},n=self["webpackChunkvue_tcptest"]=self["webpackChunkvue_tcptest"]||[];n.forEach(o.bind(null,0)),n.push=o.bind(null,n.push.bind(n))}();var n=t.O(void 0,[998],(function(){return t(9804)}));n=t.O(n)})();