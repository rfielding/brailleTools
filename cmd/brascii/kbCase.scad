$fn=12;
long = 147.3 - 0.1;
wide = 72.2 - 0.1;
high = 7.1 + 0.3;
cornerRadius = 10.1;
keybowBoard = 76.2;
ifLost = "If Lost: 703 944 7005";
email = "rob.fielding@gmail.com";
description = "8 dot braille keyboard";

module radiusBox(l,w,h,r) {
    difference() {
        scale([l, w,h]) cube(1);
        translate([0,0,0]) scale([r,r,h]) cube(1);
        translate([l-r,0,0]) scale([r,r,h]) cube(1);
        translate([0,w-r,0]) scale([r,r,h]) cube(1);
        translate([l-r,w-r,0]) scale([r,r,h]) cube(1);
    }
    translate([r,r,0]) cylinder(h=h, r1=r,r2=r);
    translate([l-r,w-r,0]) cylinder(h=h, r1=r,r2=r);
    translate([l-r,r,0]) cylinder(h=h, r1=r,r2=r);
    translate([r,w-r,0]) cylinder(h=h, r1=r,r2=r);    
}

module extrudedBox(l,w,h,r,s) {
    translate([l/2,w/2,0])
    linear_extrude(4,scale=s)
    translate([-l/2,-w/2,0])
    projection(cut=true)
    radiusBox(l,w,h,r);
}

module extrudedBox2(l,w,h,r,s) {
    translate([l/2,w/2,0])
    linear_extrude(h,scale=s)
    translate([-l/2,-w/2,0])
    projection(cut=true)
    radiusBox(l,w,h,r);
}


module flipLExtrudedBox(l,w,h,r,s) {
    translate([l/2,0,h/2])
    rotate([0,180,0])
    translate([-l/2,0,0])
    extrudedBox2(l,w,h,r,s);
}

module thePhone() {
    // main phone
    extrudedBox2(long,wide,high,cornerRadius,0.995);
    
    // touch area sticking out
    translate([0,0,high]) extrudedBox2(long,wide,high,cornerRadius,1.15);
    
    // Camera sticking out
    translate([4,wide-36.5,-high/2]) flipLExtrudedBox(36,34.1,high,6,1.5);
    
    // lbuttons sticking out
    translate([long-125,-high/2,high*0.8]) rotate([-90,0,0]) translate([0,0,0.2]) flipLExtrudedBox(43,high*0.8,high,1,1.25);
    
    // rbuttons sticking out
    translate([long-113,wide+high/2,0]) rotate([90,0,0]) translate([0,0,0.2]) flipLExtrudedBox(26,high*0.8,high,1,1.25);
    
    // bottom buttons
    difference() {
        translate([long+high/2,16,high*0.8]) rotate([0,-90,0]) rotate([0,0,90]) translate([-3,0,0.3]) flipLExtrudedBox(44+6,high*0.8,high,1,1.15);
        translate([long+high/2,0.5+wide/2-8,0]) scale([high,17,high]) cube(1);
    }
    
    // The cutout for the keyboard
    // (76-38)/2 = 19
    // (76-70)/2 = 3
    a = (keybowBoard-38)/2;
    b = (keybowBoard-70)/2;
    kbb = keybowBoard;
    translate([long-73,-2.05,-4-1]) union() {
        union() {
            translate([-5,0,0]) scale([kbb+5,kbb,2.5]) cube(1);
            translate([-5,2,2.5]) scale([kbb+5,kbb-4,2.5]) cube(1);            
            translate([-5,2,-2]) scale([kbb+5,kbb-4,6]) cube(1);
        }
    }
    
    // USB port cutout
    difference() {
        translate([long - 20/2 - 4, wide/2 - 15/2, -4.0]) scale([20,15,11]) cube(1);
    }
    
    // Round off top
    translate([-3,-3,high+1.5]) scale([long+10,wide+10,high]) cube(1);
    translate([long+3,0,high+2.2]) rotate([0,45,0]) rotate([0,90,0]) translate([0,-3,0]) scale([1,wide+6,1]) cube(1);
    translate([-3,0,high+2.2]) rotate([0,45,0]) rotate([0,90,0]) translate([0,-3,0]) scale([1,wide+6,1]) cube(1);    
}

module theCase() {
    union() {
        translate([80-7,-1.7-0.75,0]) scale([0.5, 1.1, 0.5]) rotate([90,0,0]) linear_extrude(1) text(ifLost);        
        translate([140,wide+3-1.5+1,0]) scale([-0.5, 1.1, -0.5]) rotate([-90,0,0]) linear_extrude(1) text(email);        
        translate([-3,wide,0]) scale([1,0.45,0.45]) rotate([90,0,-90]) linear_extrude(1) text(description, font="orbitreader2");
        
        translate([2.75,2.75,high+0.7]) sphere(1.5);
        translate([2.75,wide-2.75,high+0.7]) sphere(1.5);
        translate([long-2.75,2.75,high+0.7]) sphere(1.5);
        translate([long-2.75,wide-2.75,high+0.7]) sphere(1.5);
        
        translate([long-7,wide+3,-6]) scale([8,1,high+3]) cube(1);
        translate([long-7,wide+3,-6]) scale([-90,1,high-2.75]) cube(1);
        
        translate([long-7,-4,-6]) scale([8,1,high+3]) cube(1);
        translate([long-7,-4,-6]) scale([-90,1,high-2.75]) cube(1);
        
        difference() {
            translate([-3,-3,-3-4]) scale([long+6,wide+6,high+6+4]) cube(1);
            thePhone();
            // shave down corners to 45deg max
            //translate([-3,-3,-3-1.5]) rotate([45,0,0]) scale([long+6+5,2,2]) cube(1);
            //translate([-3,wide+6-3,-3-1.5]) rotate([45,0,0]) scale([long+6+5,2,2]) cube(1);
            //translate([-3,wide+6-3,-3-1.5]) rotate([0,0,-90]) rotate([45,0,0]) scale([long+6+5,2,2]) cube(1);
            //translate([long+6-3,wide+6-3,-3-1.5]) rotate([0,0,-90]) rotate([45,0,0]) scale([long+6+5,2,2]) cube(1);
            translate([-3,-3,-3]) rotate([0,0,45]) scale([1,1,40]) cube(1,center=true);
            translate([long+6-3,-3,-3]) rotate([0,0,45]) scale([1,1,40]) cube(1,center=true);
            translate([-3,wide+6-3,-3]) rotate([0,0,45]) scale([1,1,40]) cube(1,center=true);
            translate([long+6-3,wide+6-3,-3]) rotate([0,0,45]) scale([1,1,40]) cube(1,center=true);        
            
            translate([0,0,-10]) intersection() {
                translate([25,25,0]) rotate([0,0,30]) translate([-25,-25,0]) for( r = [-10:50]) {
                    for( c = [-10:50]) {
                        translate([r*5-7,c*5-7,0-1]) scale([4,4,27]) cube(1);
                    }
                }
                difference() {
                  translate([5-5,4-4,0]) scale([65/*-65+long*/+6.8,65/*-65+wide*/+6.8,25]) cube(1);
                  translate([0,wide-45,0]) scale([50,45,25]) cube(1);
                }
            }
        }
    }
}

theCase();