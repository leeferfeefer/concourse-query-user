(function(e){function r(r){for(var a,l,s=r[0],i=r[1],u=r[2],d=0,p=[];d<s.length;d++)l=s[d],o[l]&&p.push(o[l][0]),o[l]=0;for(a in i)Object.prototype.hasOwnProperty.call(i,a)&&(e[a]=i[a]);c&&c(r);while(p.length)p.shift()();return n.push.apply(n,u||[]),t()}function t(){for(var e,r=0;r<n.length;r++){for(var t=n[r],a=!0,s=1;s<t.length;s++){var i=t[s];0!==o[i]&&(a=!1)}a&&(n.splice(r--,1),e=l(l.s=t[0]))}return e}var a={},o={app:0},n=[];function l(r){if(a[r])return a[r].exports;var t=a[r]={i:r,l:!1,exports:{}};return e[r].call(t.exports,t,t.exports,l),t.l=!0,t.exports}l.m=e,l.c=a,l.d=function(e,r,t){l.o(e,r)||Object.defineProperty(e,r,{enumerable:!0,get:t})},l.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},l.t=function(e,r){if(1&r&&(e=l(e)),8&r)return e;if(4&r&&"object"===typeof e&&e&&e.__esModule)return e;var t=Object.create(null);if(l.r(t),Object.defineProperty(t,"default",{enumerable:!0,value:e}),2&r&&"string"!=typeof e)for(var a in e)l.d(t,a,function(r){return e[r]}.bind(null,a));return t},l.n=function(e){var r=e&&e.__esModule?function(){return e["default"]}:function(){return e};return l.d(r,"a",r),r},l.o=function(e,r){return Object.prototype.hasOwnProperty.call(e,r)},l.p="/";var s=window["webpackJsonp"]=window["webpackJsonp"]||[],i=s.push.bind(s);s.push=r,s=s.slice();for(var u=0;u<s.length;u++)r(s[u]);var c=i;n.push([0,"chunk-vendors"]),t()})({0:function(e,r,t){e.exports=t("56d7")},"56d7":function(e,r,t){"use strict";t.r(r);var a=t("2b0e"),o=t("bb71");t("da64");a["a"].use(o["a"],{iconfont:"md"});var n=function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("v-app",[t("v-toolbar",{attrs:{app:""}},[t("v-toolbar-title",{staticClass:"headline"},[t("span",[e._v("Bridge Troll")])])],1),t("v-content",[t("QueryUser")],1)],1)},l=[],s=function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("v-container",[t("v-layout",{attrs:{"align-center":"","justify-center":""}},[t("v-flex",{attrs:{xs12:"",sm8:"",md4:""}},[t("transition",{attrs:{name:"slide",mode:"out-in"}},[e.paid?e._e():t("v-card",{key:"toll",staticClass:"elevation-12"},[t("v-toolbar",{attrs:{dark:"",color:e.error?"red":"blue"}},[t("v-toolbar-title",[e._v(e._s(e.error?"Uh-oh!":"Enter Credentials"))])],1),t("v-card-text",[t("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(r){e.valid=r},expression:"valid"}},[t("v-text-field",{attrs:{label:"Username",outline:"",required:"",rules:e.usernameRules},model:{value:e.username,callback:function(r){e.username=r},expression:"username"}}),t("v-text-field",{attrs:{type:"password",label:"Password",outline:"",required:"",rules:e.passwordRules},model:{value:e.password,callback:function(r){e.password=r},expression:"password"}}),t("v-btn",{attrs:{disabled:!e.valid},on:{click:e.payTheTrollToll}},[e._v("\n                Pay the troll toll\n              ")]),e.error?t("label",{attrs:{color:"red"}},[e._v("Network Error - Please try again")]):e._e()],1)],1)],1),e.paid?t("v-card",{key:"success",staticClass:"elevation-12"},[t("v-toolbar",{attrs:{dark:"",color:"green"}},[t("v-toolbar-title",[e._v("Success")])],1),t("v-card-text",[e._v("\n            Toll Paid\n          ")])],1):e._e()],1)],1)],1)],1)},i=[],u=t("bc3a"),c=t.n(u),d={name:"QueryUser",props:{},data:()=>({valid:!0,paid:!1,error:!1,username:"",usernameRules:[e=>!!e||"Username is required"],password:"",passwordRules:[e=>!!e||"Password is required"]}),methods:{async payTheTrollToll(){if(this.$refs.form.validate()){const r={data:{username:btoa(this.username),password:btoa(this.password)}},t=c.a.create({baseURL:"https://localhost:3001/",timeout:1e3,headers:{"Content-Type":"application/json"}});try{await t.post("/toll",r),this.error=!1,this.paid=!0,await t.post("/shutdown")}catch(e){this.error=!0}}}}},p=d,f=(t("8e1c"),t("2877")),v=t("6544"),b=t.n(v),m=t("8336"),h=t("b0af"),y=t("99d9"),w=t("a523"),x=t("0e8f"),T=t("4bd4"),_=t("a722"),g=t("2677"),V=t("71d9"),j=t("2a7f"),k=Object(f["a"])(p,s,i,!1,null,"0e54663e",null),O=k.exports;b()(k,{VBtn:m["a"],VCard:h["a"],VCardText:y["a"],VContainer:w["a"],VFlex:x["a"],VForm:T["a"],VLayout:_["a"],VTextField:g["a"],VToolbar:V["a"],VToolbarTitle:j["a"]});var P={name:"App",components:{QueryUser:O},data:()=>({})},C=P,U=t("7496"),S=t("549c"),R=Object(f["a"])(C,n,l,!1,null,null,null),q=R.exports;b()(R,{VApp:U["a"],VContent:S["a"],VToolbar:V["a"],VToolbarTitle:j["a"]}),a["a"].config.productionTip=!1,a["a"].config.silent=!0,new a["a"]({render:e=>e(q)}).$mount("#app")},"7aed":function(e,r,t){},"8e1c":function(e,r,t){"use strict";var a=t("7aed"),o=t.n(a);o.a}});
//# sourceMappingURL=app.9b28003f.js.map