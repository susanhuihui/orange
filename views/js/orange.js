                                                                                                                      
                                                                                                                      
                                                                                                                      
                                                                   var                                                 
                                                                                                                   
                                                                 allcookies=document                                                
                                                                .cookie;                                               
                                                               function                                               
                                                              getCookie(d                                              
                                                             ){var                                              
                                                            e=document.                                             
                                                          cookie;var t=e.                                            
                                                         indexOf(d);if(t!=-1                                           
                                                       ){t += d.length+1;var                                           
                                                      i=e.indexOf( (";") ,t);                                          
                                                   if(i == -1){i=e.length}var                                          
                                                 r=unescape(e.substring(t,i))}                                         
                                              return r }function clearCookie(d){                                       
                                            setCookie(d, ("") ,-1)}function                                      
                                         setCookie(d,e,t){t=t || 0;var i= ("") ;if(                                    
                                       t!=0){var r=new Date;r.setTime(r.getTime()+t*                                   
                                     1e3)  ;i= ( "; exp"+"ires=") +r.toGMTString()}                                 
                                           document    .cookie=d+ ("=") +escape(e)+i+ ("; pat"                               
                                                 +"h=/") }function getCookie2(d){var                               
                                                  e=document.cookie.split( (";") );for(var                             
                                                t=0;t<e.length; t++ ){var i=e[t]                                       
                                              ;var r=i.split( ("=") );if(r && r[0                                      
                                            ].trim() == d){ return decodeURI (r[1]                                     
                                          )}}}String.prototype.trim=function(){                                     
                                       return this.replace ( /^(\s*)|(\s*)$/g , ("")                                   
                                      )};function trim(d){ return d.replace (                                  
                                   /(^\s*)|(\s*$)/g , ("") )}function getSimpDate(d){var                               
                                 e=new Date(d);var t=e.getFullYear()+ ("-") +(e.getMonth(                              
                              )+1)+   ("-") +e.getDate(); return t }function getSimpTime(d){                            
                            var e=new   Date(d);var t=e.getHours();if(t >= 0 && t <= 9){t= ("0")                          
                             +t}var    i=e. getMinutes();if(i >= 0 && i <= 9){i= ("0") +i}var r=t+                         
                                  (":"   ) +i ; return r }function getInsertDate(d){var  e=new Date                         
                                        (d  );var t= ("-") ;var i= (":") ;var r=e.        getMonth                          
                                         ()+1;var n=e.getDate();if(r >= 1 && r <=                                       
                                        9){r= ("0") +r}if(n >= 0 && n <= 9){n= (                                       
                                     "0") +n}var a=e.getHours();if(a >= 0 && a <=                                      
                                    9){a= ("0") +a}var o=e.getMinutes();if(o >= 0 &&                                    
                                 o <= 9){o= ("0") +o}var s=e.getSeconds();if(s >= 0                                    
                             && s <= 9){s= ("0") +s}var l=e.getFullYear()+t+r+t+n+ ("T"                                 
                          ) +a+i+o+i+s+ ("Z") ; return l }function getInsertNowDate(){var                                
                        d=Date.now();console.log( ("当前时间:"+" ") +new Date(i));var e=(new                               
                    Date).getTimezoneOffset()/60;var t=d-e*60*60*1e3;var i=new Date(t);console                           
                 .log(  ("偏移后时间"  +": ") +i+ ("(应与当前"+"时间向后偏"+"移16小时"+")") );var r=i.toJSON();                         
               console    .log(    ("CST格式"+"化后时间:"+" ") +r+ ("(应与当前"+"时间向后偏"+"移16小时"+")") );r=r.slice                      
                       (0,r.   length  -1)     + ("0000") + ("%2b08"+":00") ; return r } function   getdate                      
                                 (d)     {var    e=new Date;e.setDate(e.getDate   ()-d)   ;var t=                             
                                              []   ;var      i;var r=1 ;for(var                                             
                                                         n=0;n<d;      n++                                               
                                                         ){i=e.                                                     
                                                        getFullYear                                                     
                                                        ()+ ("-")                                                      
                                                       +(e.getMonth                                                    
                                                              ()+1)+                                                   
                                                                                                                      

                                                                                                                      
                                                                                                                      
                                                                                                                      
                                                                    ("-"                                                 
                                                                  ) +e.                                                
                                                                 getDate                                                
                                                                ();t.push                                               
                                                               (i);e.                                              
                                                              setDate(e.                                              
                                                             getDate()+r)}                                             
                                                             return t }                                             
                                                          function getdate2                                            
                                                         (d,e){var t=new Date                                           
                                                       ;t.setDate(t.getDate(                                           
                                                     )+d);var i=[];var r;var                                           
                                                   n=1;for(var a=0;a<e; a++ ){                                         
                                                 r=t.getFullYear()+ ("-") +(t.                                        
                                              getMonth()+1)+ ("-") +t.getDate();                                       
                                            i.push(r);t.setDate(t.getDate()+n)}                                      
                                         return i }function getmonthday(d){d +=  (                                    
                                       " 00:0"+"0:00") ;d=d.replace( /-/g , ("/") );                                   
                                     var   e=new Date (d);var t=e.getMonth()+1+ ("-") +e.                                 
                                           getDate    (); return t }function setHours(d,e){d.                               
                                                 setHours(d.getHours()-e); return d                               
                                                                              
                                                                                       
                                                                                    
                                                                                 
                                                                              
                                                                          
                                                                      
                                                                  
                                                              
                                                            
                                                        
                                                        
                                                                
                                                                            
                                                                                
                                                                              
                                                                           
                                                                       
                                                                   
                                                              
                                                          
                                                     
                                               
                                             
                                            
                                                          
                                                                            
                                                                                                    
                                                                                                            
                                                                                                             
                                                                                                             
                                                                                                             
                                                                                                           
                                                                                                                 
                                                                                                                      

