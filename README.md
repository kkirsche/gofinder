# GoFinder

Golang tool to parse and find specific data inside of a website's

## Installation

```
$ go get -u github.com/kkirsche/gofinder
```

## Usage

```
$ gofinder [options] [urls...]
```

## Example

```
~ ❯❯❯ gofinder -t https://www.verizon.com
INFO[0005] TITLE FOUND:	Verizon Fios & Custom TV | Internet, Cable & Phone
~ ❯❯❯ gofinder -c https://www.verizon.com
INFO[0000] COMMENT FOUND:	 VZ Tracker
INFO[0000] COMMENT FOUND:	<div class="legal content main wrap" id="disclaimer">
	*Speeds based on 75/75 Mbps. Download and upload time estimates based on maximum connection speeds. Actual speeds will vary. ^Speed and range tested in lab conditions using latest-generation wireless devices. In-home speed and range may vary depending on a number of factors like age/type of wireless devices used. Internet services, router location, distance and interference among other things. † Payment plans and prices vary based on number of connected TVs–up to 10. Media Servers and Media Clients beginning at $19.98/mo. Premium DVR Service is $32/mo. and Enhanced DVR service is $22/mo. Recorded content stored on current Set Top Box will not be transferred to new TV equipment. Specific model Verizon branded router required. FiOS Quantum TV available in select areas.</div>
INFO[0000] COMMENT FOUND:	 Modal
INFO[0000] COMMENT FOUND:	 Share
~ ❯❯❯ gofinder -l https://www.verizon.com
INFO[0000] LINK FOUND:	http://www.verizonwireless.com/?intcmp=VZ-GHOME-D-WIRELESS
INFO[0000] LINK FOUND:	/?lid=//global//residential
INFO[0000] LINK FOUND:	/home/verizonglobalhome/ghp_business.aspx
INFO[0000] LINK FOUND:	http://espanol.verizon.com/enes/home/verizonglobalhome/ghp_landing.aspx
INFO[0000] LINK FOUND:	http://www.verizonwireless.com/?intcmp=VZ-GHOME-D-WIRELESS
INFO[0000] LINK FOUND:	https://www.verizonwireless.com/smartphones/apple-iphone-x/
INFO[0000] LINK FOUND:	/?lid=//global//residential
INFO[0000] LINK FOUND:	/home/verizonglobalhome/ghp_business.aspx
INFO[0000] LINK FOUND:	/home/storelocator/
INFO[0000] LINK FOUND:	http://www.verizon.com/about/
INFO[0000] LINK FOUND:	http://www.verizon.com/about/careers/
INFO[0000] LINK FOUND:	javascript:O_LC();
INFO[0000] LINK FOUND:	http://www.verizon.com/about/privacy/
INFO[0000] LINK FOUND:	http://www.verizon.com/about/terms/
INFO[0000] LINK FOUND:	javascript:void(0);
INFO[0000] LINK FOUND:	javascript:void(0);
INFO[0000] LINK FOUND:	javascript:void(0);
~ ❯❯❯ gofinder -s https://www.verizon.com
INFO[0000] INLINE SCRIPT FOUND:	if(typeof(_satellite) == 'undefined'){
						if(document.domain.indexOf("www.verizon.com")>-1){
						document.write('<scr'+'ipt type="text/javascript" src="https://assets.adobedtm.com/10d5272d092923c410feae744225087686012423/satelliteLib-8df7d93db820b272138ecb04dbe4ed7f5023b893.js"><\/scr'+'ipt>');
						}else{
						document.write('<scr'+'ipt type="text/javascript" src="https://assets.adobedtm.com/10d5272d092923c410feae744225087686012423/satelliteLib-8df7d93db820b272138ecb04dbe4ed7f5023b893-staging.js"><\/scr'+'ipt>');
						}
					}

function receiveMessage(event){
	if(event && event != null && event.data && event.data != "null" && event.data != "") {
		var dataObj;
		try {
			if(JSON) {
				dataObj = JSON.parse(event.data);
			} else {
				dataObj = eval('(' + event.data + ')');
			}
			if(dataObj && dataObj.action && dataObj.url && (dataObj.action === "close" || dataObj.action === "order")) {
				window.location.href = dataObj.url;
			}
		} catch(e) {}
	}
}
if ( window.attachEvent ) {
	// IE
	window.attachEvent("onmessage", receiveMessage);
}
if ( document.attachEvent ) {
	// IE
	document.attachEvent("onmessage", receiveMessage);
}
if ( window.addEventListener ){
	// FF
	window.addEventListener("message", receiveMessage, false);
}
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/jquery1.11.0.min_js.js
INFO[0000] INLINE SCRIPT FOUND:	var dtmtmpurl = window.location.href.toLowerCase();
	 window.digitalData = {
            "page": {
            "sysEnv": document.domain.indexOf("www.verizon.com")!=-1?"prod":"dev",
            "adobeReportSuite": document.domain.indexOf("www.verizon.com")!=-1?"verizontelecomglobal,verizontelecomres":"verizontelecomglobaldev,verizontelecomresdev",
            "language": "en_US",
            "pfxID": dtmtmpurl.indexOf("/home")!=-1 || dtmtmpurl.indexOf("/info")!=-1 || dtmtmpurl.indexOf("/local")!=-1 ?"res":"ghp",
            "pageName": "",
            "detailPageName": "",
            "businessUnit": "residential",
            "pageType": dtmtmpurl.indexOf("/home")!=-1 || dtmtmpurl.indexOf("/info")!=-1 || dtmtmpurl.indexOf("/local")!=-1?(dtmtmpurl.indexOf("shopping") != -1?"shop":"learn"):"resghp",
            "contentChannel": dtmtmpurl,
            "lineOfBusiness": "residential",
            "cmsPlatform": "res| atlas",
            "applicationName": dtmtmpurl.indexOf("/home")!=-1 || dtmtmpurl.indexOf("/info")!=-1 || dtmtmpurl.indexOf("/local")!=-1?"residential":"global homepage",
            "events":"event3"
             },
            "userProfile": {
            "pdcString":"",
            "personalizationXP1String":"",
            "profileID":"",
            "loginStatus":"",
            "customerType":"",
            "productInterestXP1":"",
            "productQualXp1":"",
            "shoppingStageXp1":"",
            "cartContentsXp1":"",
            "existingProductsXp1":"",
            "leadOffer":"",
            "closestCableXp1":"",
            "profileIDXp1":"",
            "verizonUID":""
		    }
    }
INFO[0000] INLINE SCRIPT FOUND:	function getJSDomain() { return "//www.verizon.com"; }
					function getCSSDomain() { return "//www.verizon.com"; }
					function getImageDomain() { return "//www.verizon.com"; }
INFO[0000] INLINE SCRIPT FOUND:	var srvDomainStr = 'www.';
			function checkDomainStr(domainStr) {
				if(domainStr && (domainStr.indexOf(srvDomainStr)>-1 || domainStr.indexOf("www22")>-1))
					return true;
				else
					return false;
			}
			function getDomainStr() { return srvDomainStr; }
INFO[0000] INLINE SCRIPT FOUND:	if(window.location.hostname.indexOf('www.verizon.com') == -1){
									document.getElementById('ghpresesp').href='http://espanolstage.verizon.com/enes/home/verizonglobalhome/ghp_landing.aspx';
								}
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/businessunitcookie_js.js
INFO[0000] INLINE SCRIPT FOUND:	if(window.location.hostname.indexOf('www.verizon.com')>-1){
  s_account='verizontelecomglobal,verizontelecomres';
}else{
  s_account='verizontelecomglobaldev,verizontelecomresdev';
}
INFO[0000] EXTERNAL SCRIPT FOUND:	/includes/javascript/omnicode.js
INFO[0000] EXTERNAL SCRIPT FOUND:	/resources/verizonglobalhome/scripts/oo_engine.js
INFO[0000] INLINE SCRIPT FOUND:	vzLogging_appName="CorporateHome";
INFO[0000] EXTERNAL SCRIPT FOUND:	/resources/verizonglobalhome/scripts/pagetracker.js
INFO[0000] INLINE SCRIPT FOUND:	if(typeof s_837 != 'undefined'){
	s_837.simplepageName = "routing page";
	s_837.detailpageName = "routing page| rsp";
	s_837.pfxID = "ghp";
	s_837.prop1 = "routing page| rsp";
	s_837.prop2 = "routing";
	s_837.prop3 = "routing";
	s_837.prop4 = "/vz/home/residential";
	s_837.prop6 = "routing";
	s_837.prop48 = "global homepage";
}

//var s_code=s_837.t();
//if(s_code)document.write(s_code);
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/vzr.aims_chat_js.js
INFO[0000] INLINE SCRIPT FOUND:	document.write(new Date().getFullYear())
INFO[0000] INLINE SCRIPT FOUND:	var bubbleClicks = function(){
	var lob = document.getElementsByClassName('cmb-routing-text');
	for(var i = 0; i < lob.length; i++){
		//console.log(lob[i].getElementsByTagName('a')[0].getAttribute('href'));
		lob[i].getElementsByTagName('a')[0].onclick = function(event){
			event.preventDefault();
			//window.location.href = this.getElementsByTagName('a')[0].getAttribute('href');
		}
	}
}
bubbleClicks();

var scbmfblink=document.getElementById('scbmfeedbackghp');
if(scbmfblink!=null){
	scbmfblink.href=scbmfblink.href+'&referer='+encodeURIComponent(window.location.href);
}
INFO[0000] INLINE SCRIPT FOUND:	var master_cookie_on = document.getElementById("mastercookieon").value;
if(document.getElementById("currentDomain").value==""){
	document.getElementById("currentDomain").value='http://' + document.domain;
}
if(document.getElementById("currentDomainVZ").value==""){
	document.getElementById("currentDomainVZ").value='http://' + document.domain;
}
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/vzcom_js.js
INFO[0000] EXTERNAL SCRIPT FOUND:	/home/ak-cached/2h/javascript/commons.js
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/sitecatalyst_footer_m_js.js
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/ctahandled_m_js.js
INFO[0000] EXTERNAL SCRIPT FOUND:	/cs/groups/public/documents/adacct/xp1commonjs.js
INFO[0000] INLINE SCRIPT FOUND:	$( document ).ready(function() {
    setTimeout('makeXP1Call()',3500);
	 setCustomValues();
});
INFO[0000] INLINE SCRIPT FOUND:	if(typeof(_satellite)!='undefined'){
_satellite.pageBottom();
}
function setCustomValues(){
 if(document.getElementById('xp1ProfileId')){
 var cstmxp1 = document.getElementById('xp1ProfileId').value;
     if(cstmxp1.indexOf("v36")>-1){
	  pdcupdate("/home/gsmservlet?Mode=write&vzpers=custv-A");
	 }else if(cstmxp1.indexOf("v35")>-1){
	  pdcupdate("/home/gsmservlet?Mode=write&vzpers=custv-S");
	 }else if(cstmxp1.indexOf("v521_002")>-1){
	  pdcupdate("/home/gsmservlet?Mode=write&VzOSEL=SAPRO-LOCK");
	 }else if(cstmxp1.indexOf("v521_001")>-1){
	  pdcupdate("/home/gsmservlet?Mode=write&VzOSEL=SAPRO-STAN");
	 }
  }
}
```
