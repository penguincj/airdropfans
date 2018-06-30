
 
 $(".navbar-header").click(function(){
  
    $(".glyphicon-align-justify").toggle();
    $(".glyphicon-remove").toggle();
      $(".glyphicon-one").show();
    $(".glyphicon-two").hide();
    $(".glyphicon-five").show();
    $(".glyphicon-six").hide();
    $(".glyphicon-three").show();
    $(".glyphicon-four").hide();
  
  });
  
  $(".li2").click(function(){
  
    $(".glyphicon-one").toggle();
    $(".glyphicon-two").toggle();
    $(".glyphicon-five").show();
    $(".glyphicon-six").hide();
    $(".glyphicon-three").show();
    $(".glyphicon-four").hide();
   
  });

  $(".li3").click(function(){
  
    $(".glyphicon-three").toggle();
    $(".glyphicon-four").toggle();
    $(".glyphicon-one").show();
    $(".glyphicon-two").hide();
    $(".glyphicon-five").show();
    $(".glyphicon-six").hide();
  
  });
  $(".li1").click(function(){
  
    $(".glyphicon-five").toggle();
    $(".glyphicon-six").toggle();  
    $(".glyphicon-three").show();
    $(".glyphicon-four").hide();
    $(".glyphicon-one").show();
    $(".glyphicon-two").hide();
   
  });

 

$(document).ready(function(){
    $(".prduct-on").mouseover(function(){
         
        $(".connetos-on").animate({
            left:'4%'      
        },200);
      });
     $(".prduct-on").mouseout(function(){
         
        $(".connetos-on").animate({
            left:'0'           
        },10);
      });
  });
$(document).ready(function(){
    $(".width-full-1").mouseover(function(){
         
        $(".row-text-2").animate({
            opacity:'0.5'           
        },100);
        $(".row-text-3").animate({
            opacity:'0.5'           
        },100);      
    });
    $(".width-full-1").mouseout(function(){
         
        $(".row-text-2").animate({
            opacity:'1'           
        },100);
        $(".row-text-3").animate({
            opacity:'1'           
        },100);   
    });
});
 
$(document).ready(function(){
    $(".width-full-2").mouseover(function(){
        
        $(".row-text-1").animate({
            opacity:'0.5'           
        },100);
        $(".row-text-3").animate({
            opacity:'0.5'           
    
        },100);

    });
    $(".width-full-2").mouseout(function(){
         
        $(".row-text-1").animate({
            opacity:'1'           
        },100);
        $(".row-text-3").animate({
            opacity:'1'           
        },100);
    });
});
$(document).ready(function(){
    $(".width-full-3").mouseover(function(){
         
        $(".row-text-1").animate({
            opacity:'0.5'           
        },100); 
        $(".row-text-2").animate({
            opacity:'0.5'           
        },100);
    });
    $(".width-full-3").mouseout(function(){
         
        $(".row-text-1").animate({
            opacity:'1'           
        },100);
        $(".row-text-2").animate({
            opacity:'1'           
        },100); 
    });
});


$(document).ready(function(){
  $(".list-group-item-1").click(function(){
    $(".media-1").animate({
      opacity:'0'
    },10);
     $(".media-2").animate({
      opacity:'0'
    },10);
      $(".media-3").animate({
      opacity:'0'
    },10);
       $(".media-4").animate({
      
      top:'-527px'
    },10);
  });
  $(".list-group-item-2").click(function(){
     $(".media-2").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-3").click(function(){
     $(".media-1").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-4").click(function(){
     $(".media-3").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-0").click(function(){
    $(".media-1").animate({
      opacity:'1',
      top:'0px'
    },10);
    
    $(".media-2").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-3").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-4").animate({
      opacity:'1',
       top:'0px'
    },10);
  });
});

$(document).ready(function(){
  $(".list-group-item-2").click(function(){
    $(".media-1").animate({
      opacity:'0'
    },10);
     $(".media-3").animate({
      opacity:'0'
    },10);
      $(".media-4").animate({
      opacity:'0'
    },10);
       $(".media-2").animate({
      
      top:'-200px'
    },10);
  });
    $(".list-group-item-1").click(function(){
     $(".media-4").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-3").click(function(){
     $(".media-1").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-4").click(function(){
     $(".media-3").animate({
      opacity:'1'
    },10);
   });
   $(".list-group-item-0").click(function(){
    $(".media-1").animate({
      opacity:'1',
      top:'0px'
    },10);
    
    $(".media-2").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-3").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-4").animate({
      opacity:'1',
       top:'0px'
    },10);
  });
});

$(document).ready(function(){
  $(".list-group-item-3").click(function(){
    $(".media-2").animate({
      opacity:'0'
    },10);
     $(".media-3").animate({
      opacity:'0'
    },10);
      $(".media-4").animate({
      opacity:'0'
    },10);
       $(".media-4").animate({
      
      top:'0px'
    },10);
  });
   $(".list-group-item-2").click(function(){
     $(".media-2").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-1").click(function(){
     $(".media-4").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-4").click(function(){
     $(".media-3").animate({
      opacity:'1'
    },10);
   });
   $(".list-group-item-0").click(function(){
    $(".media-1").animate({
      opacity:'1',
      top:'0px'
    },10);
    
    $(".media-2").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-3").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-4").animate({
      opacity:'1',
       top:'0px'
    },10);
  });
});

$(document).ready(function(){
  $(".list-group-item-4").click(function(){
    $(".media-1").animate({
      opacity:'0'
    },10);
     $(".media-2").animate({
      opacity:'0'
    },10);
      $(".media-4").animate({
      opacity:'0'
    },10);
       $(".media-3").animate({
      
      top:'-360px'
    },10);
  });
   $(".list-group-item-2").click(function(){
     $(".media-2").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-3").click(function(){
     $(".media-1").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-1").click(function(){
     $(".media-4").animate({
      opacity:'1'
    },10);
   });
  $(".list-group-item-0").click(function(){
    $(".media-1").animate({
      opacity:'1',
      top:'0px'
    },10);
    
    $(".media-2").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-3").animate({
      opacity:'1',
       top:'0px'
    },10);
    
    $(".media-4").animate({
      opacity:'1',
       top:'0px'
    },10);
  });
});
/*
function check(form) {
          if(form.userId.value=='') {
            document.getElementById('name').placeholder="请输入您的名字";
            document.getElementById('email').placeholder="请输入正确格式的Email";
            document.getElementById('phone').placeholder="请输入您的电话号码";
                form.userId.focus();
                return false;
           }
       if(form.email.value==''){
                document.getElementById('email').placeholder="请输入正确格式的Email";
                document.getElementById('phone').placeholder="请输入您的电话号码";
                form.email.focus();
                return false;
         }
          if(form.phone_number.value==''){
                document.getElementById('phone').placeholder="请输入您的电话号码";
                form.phone_number.focus();
                return false;
         }
        if(form.phone_number.value.length<7||form.phone_number.value.length>11){
                  alert("您输入的电话号码不正确");
                  form.phone_number.focus();
                  return false;
           }
  
          if(form.value!==''){
                alert("恭喜您，提交成功!");
                form.phone_number.focus();
                return true;
         }
         
         return true;
         }
         */
         var app = angular.module("myNoteApp", []);
         app.controller("myNoteCtrl", function($scope) {
            $scope.message = "";
            $scope.left  = function() {return 250 - $scope.message.length;};
            $scope.clear = function() {$scope.message = "";};
            $scope.save  = function() {alert("Note Saved");};
     
          });