// JavaScript Document<script type="text/javascript">
	function changeStyle(){
	 this.onclick=function(){
	 var list=this.parentNode.childNodes;
	 for(var i=0;i<list.length;i++){
	  if(1==list[i].nodeType && 'active'==list[i].className){
		list[i].className="";
	  }
	 }
	 this.className='active';
	 }
	}
	 var tabs=document.getElementById('tabnav').childNodes;
	 for(var i=0;i<tabs.length;i++){
	 if(1==tabs[i].nodeType){
	  changeStyle.call(tabs[i]);
	 }
	};
	
	
	
	function qiehuantap(tapnum){
							for(var i =1;i<=3;i++){
								document.getElementById("divtap"+i).style.display="none";
								document.getElementById("limenu"+i).className="";
							}
							document.getElementById("divtap"+tapnum).style.display="block";
								document.getElementById("limenu"+tapnum).className="active";
							};
                        function jilutap(tapnum){
							for(var i =1;i<=2;i++){
								document.getElementById("xfjl"+i).style.display="none";
								document.getElementById("tapli"+i).className="";
							}
							document.getElementById("xfjl"+tapnum).style.display="block";
								document.getElementById("tapli"+tapnum).className="active";
							};
	
	function openShutManager(oSourceObj,oTargetObj,shutAble,oOpenTip,oShutTip){
			var sourceObj = typeof oSourceObj == "string" ? document.getElementById(oSourceObj) : oSourceObj;
			var targetObj = typeof oTargetObj == "string" ? document.getElementById(oTargetObj) : oTargetObj;
			var openTip = oOpenTip || "";
			var shutTip = oShutTip || "";
			if(targetObj.style.display!="none"){
			   if(shutAble) return;
			   targetObj.style.display="none";
			   if(openTip && shutTip){
				sourceObj.innerHTML = shutTip; 
			   }
			} else {
			   targetObj.style.display="block";
			   if(openTip && shutTip){
				sourceObj.innerHTML = openTip; 
			   }
			}
			};
			
			
			
			function openShutManagers(oSourceObj,oTargetObj,shutAble,oOpenTip,oShutTip){
			var sourceObj = typeof oSourceObj == "string" ? document.getElementById(oSourceObj) : oSourceObj;
			var targetObj = typeof oTargetObj == "string" ? document.getElementById(oTargetObj) : oTargetObj;
			var openTip = oOpenTip || "";
			var shutTip = oShutTip || "";
			if(targetObj.style.display!="none"){
			   if(shutAble) return;
			   targetObj.style.display="none";
			   if(openTip && shutTip){
				sourceObj.innerHTML = shutTip; 
			   }
			} else {
			   targetObj.style.display="block";
			   if(openTip && shutTip){
				sourceObj.innerHTML = openTip; 
			   }
			}
			};