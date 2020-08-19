package livejs

var js = `!function(){var e={Etag:1,"Last-Modified":1,"Content-Length":1,"Content-Type":1},t={},o={},a={},n={},s=!1,r={html:1,css:1,js:1},c={heartbeat:function(){document.body&&(s||c.loadresources(),c.checkForChanges()),setTimeout(c.heartbeat,1e3)},loadresources:function(){function e(e){var t=document.location,o=new RegExp("^\\.|^/(?!/)|^[\\w]((?!://).)*$|"+t.protocol+"//"+t.host);return e.match(o)}for(var o=document.getElementsByTagName("script"),n=document.getElementsByTagName("link"),i=[],l=0;l<o.length;l++){var d=o[l].getAttribute("src");if(d&&e(d)&&i.push(d),d&&d.match(/\blive.js#/)){for(var u in r)r[u]=null!=d.match("[#,|]"+u);d.match("notify")&&alert("Live.js is loaded.")}}r.js||(i=[]),r.html&&i.push(document.location.href);for(l=0;l<n.length&&r.css;l++){var h=n[l],m=h.getAttribute("rel"),p=h.getAttribute("href",2);p&&m&&m.match(new RegExp("stylesheet","i"))&&e(p)&&(i.push(p),a[p]=h)}for(l=0;l<i.length;l++){var f=i[l];c.getHead(f,function(e,o){t[e]=o})}var v=document.getElementsByTagName("head")[0],g=document.createElement("style"),w="transition: all .3s ease-out;";css=[".livejs-loading * { ",w," -webkit-",w,"-moz-",w,"-o-",w,"}"].join(""),g.setAttribute("type","text/css"),v.appendChild(g),g.styleSheet?g.styleSheet.cssText=css:g.appendChild(document.createTextNode(css)),s=!0},checkForChanges:function(){for(var e in t)o[e]||c.getHead(e,function(e,o){var a=t[e],n=!1;for(var s in t[e]=o,a){var r=a[s],i=o[s],l=o["Content-Type"];switch(s.toLowerCase()){case"etag":if(!i)break;default:n=r!=i}if(n){c.refreshResource(e,l);break}}})},refreshResource:function(e,t){switch(t.toLowerCase()){case"text/css":var o=a[e],s=document.body.parentNode,r=o.parentNode,i=o.nextSibling,l=document.createElement("link");s.className=s.className.replace(/\s*livejs\-loading/gi,"")+" livejs-loading",l.setAttribute("type","text/css"),l.setAttribute("rel","stylesheet"),l.setAttribute("href",e+"?now="+1*new Date),i?r.insertBefore(l,i):r.appendChild(l),a[e]=l,n[e]=o,c.removeoldLinkElements();break;case"text/html":if(e!=document.location.href)return;case"text/javascript":case"application/javascript":case"application/x-javascript":document.location.reload()}},removeoldLinkElements:function(){var e=0;for(var t in n){try{var o=a[t],s=n[t],r=document.body.parentNode,i=o.sheet||o.styleSheet;(i.rules||i.cssRules).length>=0&&(s.parentNode.removeChild(s),delete n[t],setTimeout(function(){r.className=r.className.replace(/\s*livejs\-loading/gi,"")},100))}catch(t){e++}e&&setTimeout(c.removeoldLinkElements,50)}},getHead:function(t,a){o[t]=!0;var n=window.XMLHttpRequest?new XMLHttpRequest:new ActiveXObject("Microsoft.XmlHttp");n.open("HEAD",t,!0),n.onreadystatechange=function(){if(delete o[t],4==n.readyState&&304!=n.status){n.getAllResponseHeaders();var s={};for(var r in e){var c=n.getResponseHeader(r);"etag"==r.toLowerCase()&&c&&(c=c.replace(/^W\//,"")),"content-type"==r.toLowerCase()&&c&&(c=c.replace(/^(.*?);.*?$/i,"$1")),s[r]=c}a(t,s)}},n.send()}};"file:"!=document.location.protocol?(window.liveJsLoaded||c.heartbeat(),window.liveJsLoaded=!0):window.console&&console.log("Live.js doesn't support the file protocol. It needs http.")}();`