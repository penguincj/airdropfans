

 (function($){
	$(document).ready(function(){
	
	 
		
		
	    	// Fixed header
		//-----------------------------------------------
		$(window).scroll(function() {
        	if (($(".fixed").length > 0)) { 
	    	    if(($(this).scrollTop() > 0) && ($(window).width() > 767)) {
					$("body").addClass("fixed-header-on");
				} else {
					$("body").removeClass("fixed-header-on");
				}
			};
		});

		$(window).load(function() {
			if (($(".fixed").length > 0)) { 
				if(($(this).scrollTop() > 0) && ($(window).width() > 767)) {
					$("body").addClass("fixed-header-on");
				} else {
					$("body").removeClass("fixed-header-on");
				}
			};
		});

		//Scroll Spy
		//-----------------------------------------------
		if($(".scrollspy").length>0) {
			$("body").addClass("scroll-spy");
			$('body').scrollspy({ 
				target: '.scrollspy',
				offset: 152
			});
		}

		//Smooth Scroll
		//-----------------------------------------------
		if ($(".smooth-scroll").length>0) {
			$('.smooth-scroll a[href*=#]:not([href=#]), a[href*=#]:not([href=#]).smooth-scroll').click(function() {
				if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'') && location.hostname == this.hostname) {
					var target = $(this.hash);
					target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
					if (target.length) {
						$('html,body').animate({
							scrollTop: target.offset().top-151
						}, 1000);
						return false;
					}
				}
			});
		}

		// Animations
		//-----------------------------------------------
		if (($("[data-animation-effect]").length>0) && !Modernizr.touch) {
			$("[data-animation-effect]").each(function() {
				var $this = $(this),
				animationEffect = $this.attr("data-animation-effect");
				if(Modernizr.mq('only all and (min-width: 768px)') && Modernizr.csstransitions) {
					$this.appear(function() {
						setTimeout(function() {
							$this.addClass('animated object-visible ' + animationEffect);
						}, 400);
					}, {accX: 0, accY: -130});
				} else {
					$this.addClass('object-visible');
				}
			});
		};

		// Isotope filters
		//-----------------------------------------------
		if ($('.isotope-container').length>0) {
			$(window).load(function() {
				$('.isotope-container').fadeIn();
				var $container = $('.isotope-container').isotope({
					itemSelector: '.isotope-item',
					layoutMode: 'masonry',
					transitionDuration: '1s',
					filter: "*"
				});
				// filter items on button click
				$('.filters').on( 'click', 'p a', function() {
					var filterValue = $(this).attr('data-filter');
					$(".filters").find("li.active").removeClass("active");
					$(this).parent().addClass("active");
					$container.isotope({ filter: filterValue });
					return false;
				});
			});
		};

		//Modal
		//-----------------------------------------------
		if($(".modal").length>0) {
			$(".modal").each(function() {
				$(".modal").prependTo( "body" );
			});
		}

	}); // End document ready
})(this.jQuery);


  
 $(document).ready(function(){
  $(".image4").mouseenter(function(){

    $(".image4 img").attr("src","images/chilun.gif");
  });
  $(".image4").mouseleave(function(){

    $(".image4 img").attr("src","images/chilun_tu.jpg");
  });
});

 $(document).ready(function(){
  $(".image2").mouseenter(function(){

    $(".image2 img").attr("src","images/yunduo.gif");
  });
  $(".image2").mouseleave(function(){

    $(".image2 img").attr("src","images/yunduo_tu.jpg");
  });
});
 
$(document).ready(function(){
  $(".image3").mouseenter(function(){

    $(".image3 img").attr("src","images/caozuotai.gif");
  });
  $(".image3").mouseleave(function(){

    $(".image3 img").attr("src","images/caozuotai_tu.png");
  });
});
   
$(function(){		 		 
    $("#carousele").carousel('cycle');       
});

$(function(){		 		 
    $("#myCarousel").carousel('cycle');       
});

$(function(){        
    $("#carousel").carousel('cycle');
});

$(function(){         
    $(".o3").carousel('cycle');
});

$(document).ready(function(){
    $(".bg-image-1-image").backstretch('../images/network10.jpg');
 });
 
$(document).ready(function(){
    $(".bg-image-2-image").backstretch('../images/bg-image-2.jpg');
 });

$(document).ready(function(){
    $(".bg-image-3-image").backstretch('../images/banner3.jpg');

});

$(document).ready(function(){
    $(".bg-image-5-image").backstretch('../images/portfolio-12.jpg');

});

$(document).ready(function(){
    $(".bg-image-about-image").backstretch('../images/about/surfers.jpg');

});
 
$(document).ready(function(){
    $(".bg-image-on-image").backstretch('../images/neuro.jpg');

});

$(document).ready(function(){
    $(".bg-image-connetos-express-image").backstretch('../images/products-connetos-express-bg.jpg');

});

/*
$(document).ready(function(){  
    $(".li1").click(function(){
        $(".ul1").show();
    });
    $(".li1").click(function(){
        $(".ul1").hide();   
    });
});
$(document).ready(function(){  
    $(".li2").click(function(){
        $(".ul2").show();
    });
    $(".li2").click(function(){
        $(".ul2").hide();   
    });
});
$(document).ready(function(){  
    $(".li3").click(function(){
        $(".ul3").show();
    });
    $(".li3").click(function(){
        $(".ul3").hide();   
    });
});
*/
$(document).ready(function(){  
    $(".btn-info").click(function(){
        $(".uli").slideDown();
    });
    $(".btn-info").click(function(){
        $(".uli").slideUp();   
    });
});
$(document).ready(function(){
  $(".address").click(function(){
    $(".bad").slideToggle("slow");
  });  
});

$(document).ready(function(){
    $("#goToTop").hide()//隐藏go to top按钮
        $(function(){
            $(window).scroll(function(){
                if($(this).scrollTop()>7){//当window的scrolltop距离大于1时，go to top按钮淡出，反之淡入
                    $("#goToTop").fadeIn();
                } else {
                    $("#goToTop").fadeOut();
                }
            });
        });
     

        // 给go to top按钮一个点击事件
$("#goToTop a").click(function(){
    $("html,body").animate({scrollTop:0},400);//点击go to top按钮时，以800的速度回到顶部，这里的800可以根据你的需求修改
            return false;
        });
    });


 

$(document).ready(function(){
    $(".footerweixin").mouseenter(function(){
    $(".dimension").show();
   });
    $(".dimension").mouseenter(function(){
    $(".dimension").show();
   });
  $(".footerweixin").mouseleave(function(){
    $(".dimension").hide(); 
  });
  $(".dimension").mouseleave(function(){
    $(".dimension").hide(); 
  });
});
function heredoc(fn) {
    return fn.toString().split('\n').slice(1,-1).join('\n') + '\n'
}
content = heredoc(function(){/*
admin@host-b:/$ ssh root@192.168.1.31<br>
root@192.168.1.31‘s password<br>
 
Last login: Wed Mar 29 15:12:01 2017 from 192.168.1.101<br>
    ______                            __   ____  _____<br>
   / ____/____   ____   ____   ___   / /_ / __ \/ ___/<br>
  / /    / __ \ / __ \ / __ \ / _ \ / __// / / /\__ \ <br>
 / /___ / /_/ // / / // / / //  __// /_ / /_/ /___/ / <br> 
 \____/ \____//_/ /_//_/ /_/ \___/ \__/ \____//____/  <br> 

root@BJ-YUNQI-C1020-31.Int:~ 1# <br>
root@BJ-YUNQI-C1020-31.Int:~ 1# cd /<br>
root@BJ-YUNQI-C1020-31.Int:/ 1# ls<br>
bin  boot  dev  etc  hello.c  home  lib  lib64  root  sbin  switch  sys  tmp  usr  var<br>
root@BJ-YUNQI-C1020-31.Int:/ 1# cat hello.c<br>
#include "stdio.h"<br>
int main(void)<br>
{<br>
    printf(“Hello World on ConnetOS!\n”);<br>
    return 0;<br>
}<br>
root@BJ-YUNQI-C1020-31.Int:/ 1# gcc -o hello hello.c<br>
root@BJ-YUNQI-C1020-31.Int:/ 1# ./hello<br>
Hello World on ConnetOS!<br>
root@BJ-YUNQI-C1020-31.Int:/ 1#<br>
root@BJ-YUNQI-C1020-31.Int:/ 1# cli<br>
Welcome to switch CLI shell.<br>
BJ-YUNQI-C1020-31.Int 1> show arp<br>
Aging-time(seconds): 1200<br>
Total count        : 6<br>
Address          HW Address         Type     Interface  Age<br>
--------------- -----------------  -------  ---------  -----<br>
7.7.7.7          00:00:00:11:22:77  Static   vlan7      N/A<br>
9.9.9.9          00:00:00:11:22:99  Static   vlan9      N/A<br>
11.11.11.10      00:00:00:AA:AA:44  Dynamic  vlan100    102<br>
11.11.11.20      00:00:00:BB:BB:22  Dynamic  vlan100    156<br>
22.22.22.20      2C:60:0C:7B:C1:FB  Dynamic  vlan20     229<br>
33.33.33.20      2C:60:0C:7B:C1:FB  Dynamic  vlan30     229<br>
BJ-YUNQI-C1020-31.Int 1> <br>
BJ-YUNQI-C1020-31.Int 1> exit<br>
root@BJ-YUNQI-C1020-31.Int:/ 1# <br>
root@BJ-YUNQI-C1020-31.Int:/ 1# exit<br>
Logout<br>
Connection to 192.168.1.31 closed.
*/});
 i = 0;
 function show(){  
 str  = content.substr(0,i);
 txt.innerHTML = str + "_"; 
 i++;   
if (i>content.length)i=0; 
 setTimeout("show()",10); }

  
 






  



 

 
