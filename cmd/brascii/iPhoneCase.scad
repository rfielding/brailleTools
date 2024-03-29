// iPhone case. I measured in cm, so mult units times 10
long=14.725 + 0.03 + 0.02;
wide=7.2 + 0.02;
high=0.71;
t=0.25;
lbuttonsStart=8.6;
lbuttonsStop=12.4+0.1;
rbuttonsStart=9;
rbuttonsStop=11;
bStart=1.6;
bStop=6;
camLongStart=11.0;
camLongStop=14.4;
camWideStart=0.4;
camWideStop=3.65;
corner=0.58;
btnSize=1.25;
btnDist=0.4;
logoCenter=7.4;
logoRadius=1.5;

/*
owner="Richard Fielding";
reward="If Found: 703 789 3318";
email="windwalkers@gmail.com";
*/

///*
owner="Rob Fielding";
reward="If Found: 703 944 7005";
email="rob.fielding@gmail.com";
// */


/*
owner="Danica Fielding";
reward="If Found: 703 944 7005";
email="danica.fielding77@icloud.com";
*/

if(owner == "Richard Fielding") {
    reward="If Found: 703 789 3318";
    email="windwalkers@gmail.com";
}



module thePhone() {
scale([10,10,10])
difference() {
  union() {
    translate([0.25-long/2,-wide/2,-0.15]) rotate([90,0,0]) scale([0.045*0.82,0.045,0.045]) linear_extrude(4) text(email, font="monospace");    
    translate([8.5-long/2,wide/2,-0.15]) rotate([90,0,180]) scale([0.045,0.045,0.045]) linear_extrude(4) text(reward, font="monospace");
    translate([long/2,1.5-wide/2-0.7,-0.15]) rotate([90,0,90]) scale(0.045) linear_extrude(4) text(owner, font="monospace");
    translate([-long/2 + lbuttonsStop - 0.8-0.2,wide/2-0.008,-high-0.19]) scale([0.2,t/2+0.001,high+0.5]) cube(1);
    difference() {      
      translate([long/2-corner+t/2,wide/2-corner+t/2,-high/2]) scale([corner,corner,high+t]) cube(1);
      translate([long/2-corner+t/2,wide/2-corner+t/2-0.4,-high/2]) rotate([0,0,45]) scale([corner,corner,high+0.025]) cube(1);
    }
    difference() {
      translate([-(long/2-corner+t/2),wide/2-corner+t/2,-high/2]) scale([-corner,corner,high+t]) cube(1);    
      translate([-(long/2-corner+t/2)+corner-t/2,wide/2-corner+t/2,-high/2]) rotate([0,0,45]) scale([-corner,corner,high+0.025]) cube(1);             
    }
    difference() {
      translate([(long/2-corner+t/2),-(wide/2-corner+t/2),-high/2]) scale([corner,-corner,high+t]) cube(1);
      translate([(long/2-corner+t/2)-0.4,-(wide/2-corner+t/2),-high/2]) rotate([0,0,45]) scale([corner,-corner,high+0.025]) cube(1);        
    }
    difference() {    
      translate([-(long/2-corner+t/2),-(wide/2-corner+t/2),-high/2]) scale([-corner,-corner,high+t]) cube(1);            
      // This corner always breaks because supports won't come off completely                                                                                   
      translate([-(long/2-corner+t/2)+0.4,-(wide/2-corner+t/2),-high/2]) rotate([0,0,-45]) scale([-corner,-corner,high+0.025]) cube(1);            
    }
    difference() {
      union() {
        translate([0,0,-0.025]) scale([long+t, wide+t, high+t+0.045]) cube(1,center=true);
      }
      union() {
        translate([-long/2-0.98,wide/2-bStart,-4.7]) scale([1,-1*(bStop-bStart),5]) cube(1);
        translate([camLongStart - long/2,-(wide/2-camWideStart),-1]) scale([(camLongStop-camLongStart),(camWideStop-camWideStart),1]) cube(1);
        translate([rbuttonsStart-long/2,-3.6,-4.7])   scale([rbuttonsStop-rbuttonsStart,-5,5]) cube(1);
        translate([lbuttonsStart-long/2,3.6,-4.7])  scale([lbuttonsStop-lbuttonsStart,5,5]) cube(1);
        // The actual gap for the iPhone!
        scale([long,wide,high]) cube(1,center=true);
        translate([0,0,high/2]) scale([long,wide,high]) cube(1,center=true);
      }
    }  
  }
  union() {
    translate([long/2+2*t,wide/2-2*t,-high]) rotate([0,0,45]) scale([1,1,1]) cube(2);
    translate([long/2+2*t,-(wide/2-2*t),-high]) rotate([0,0,45]) scale([-1,-1,1]) cube(2);      
    translate([-(long/2+2*t),wide/2-2*t,-high]) rotate([0,0,45]) scale([1,1,1]) cube(2);
    translate([-(long/2+2*t),-(wide/2-2*t),-high]) rotate([0,0,45]) scale([-1,-1,1]) cube(2);            
      
    translate([long/2+2*t,0,-high]) rotate([0,45,0]) scale([1,10,1]) cube(1,center=true);
    translate([-(long/2+2*t),0,-high]) rotate([0,45,0]) scale([1,10,1]) cube(1,center=true);
    translate([0,(wide/2+2*t),-high]) rotate([45,0,0]) scale([20,1,1]) cube(1,center=true);
    translate([0,-(wide/2+2*t),-high]) rotate([45,0,0]) scale([20,1,1]) cube(1,center=true);      
      
    translate([long/2+2*t,0,high]) rotate([0,45,0]) scale([1,10,1]) cube(1,center=true);
    translate([-(long/2+2*t),0,high]) rotate([0,45,0]) scale([1,10,1]) cube(1,center=true);
    translate([0,(wide/2+2*t),high]) rotate([45,0,0]) scale([20,1,1]) cube(1,center=true);
    translate([0,-(wide/2+2*t),high]) rotate([45,0,0]) scale([20,1,1]) cube(1,center=true);

    translate([0,0,0]) difference() {
        translate([camLongStart - long/2,-(wide/2-camWideStart),0]) scale([(camLongStop-camLongStart),(camWideStop-camWideStart),1]) cube(1);
        union() {
            translate([camLongStart - long/2 + (camLongStop-camLongStart),-(wide/2-camWideStart),0]) rotate([0,-45,0]) scale([(camLongStop-camLongStart),(camWideStop-camWideStart),2]) cube(1);
            translate([camLongStart - long/2 ,-(wide/2-camWideStart),0]) rotate([0,45,0]) scale([-(camLongStop-camLongStart),(camWideStop-camWideStart),2]) cube(1);   
            translate([camLongStart - long/2,-(wide/2-camWideStart),0]) rotate([45,0,0]) scale([(camLongStop-camLongStart),(camWideStop-camWideStart),1]) cube(1);
            translate([camLongStart - long/2,-(wide/2-camWideStart)+(camWideStop-camWideStart),0]) rotate([-45,0,0]) scale([(camLongStop-camLongStart),-(camWideStop-camWideStart),1]) cube(1);            
        }
    }
    
    translate([camLongStart-long/2-0.155,camWideStart-wide/2-0.18,-high/2-t*0.7]) scale([(camLongStop-camLongStart),(camWideStop-camWideStart),(camWideStop-camWideStart)])
    scale([1.1,1.1,1.1]) difference() {
      scale([1,1,0.7]) cube(1);
      union() {
          translate([1,0,0]) rotate([0,-45,0]) cube(1);
          translate([0,0,0]) rotate([0,-45,0]) cube(1);
          translate([0,1,0]) rotate([45,0,0]) cube(1);     
          translate([0,0,0]) rotate([45,0,0]) cube(1);                    
      }
    }
  }
}
}


if(owner == "Rob Fielding") {
    difference() {
        thePhone();
        //translate([-long/2,wide/2-3.5,-6]) linear_extrude(height=2) scale([0.4,0.4,0.1]) import("logo.svg",center=true);
        translate([-long/2-21,-wide/2+3.5,-6]) scale([76,76,4]) cube(1,center=true);
        translate([-long/2-60,-wide/2+4,-6]) scale([15,10,4]) cube(1,center=true);
    }
}

if(owner == "Richard Fielding") {
    difference() {
        thePhone();
        translate([-long/2,wide/2-4.5,-4.5]) scale(0.35) rotate([0,180,-90]) linear_extrude(3) import("usmc.svg",center=true);
    }
}

if(owner == "Danica Fielding") {
    difference() {
        thePhone();
        difference() {
            translate([-long/2,wide/2-1.5,-4.5]) scale(0.5) rotate([0,180,-90]) linear_extrude(3) import("plants.svg", center=true);
            translate([30,-35,-10]) scale([40,40,10]) cube(1);
            translate([-120,-50,-10]) scale([50,150,50]) cube(1);
        }
    }
}