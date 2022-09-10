//// parameters of cell
// Note that the primary challenge for making braille
// templates is in taking the geometry of the stylus into account
//
// You generally want to re-use the same stylus for many fonts.
// Flex the dimensions around it, while respecting the stylus.
// This means that measurements from center to center are used.
// The dome at the end of tip stylus must be taken into account.
// It is usually not a half-sphere, but a flattened half-sphere.
// Most paper will rip if it is a half-spere; with every dot
// making a small hole that feels different in texture.


// Set to this either 6 or 8, for the kind of braille
dots=8;

// standard spacing between dot centers is 2.5. As low as 2.0 works,
// which is the stylus diameter
dot2dotH=2.3;
dot2dotV=2.3;

// extra spacing between lines (in addition to between dots
lineExtra=0.4;

// spacing between cells at same dot
cell2cell=6.0;

// the diameter of the stylus
// this is very important to get proper domes
// if it's too large, then dots will be in wrong places.
stylusDiameter=2;

// A value of 1.0 is a perfectly round dome.
// A perfectly round dome leaves rather large holes with paper rips
// A value of 0.75 gives a slightly flat ellipsoid pit
dome=0.75;

// parameters of slate in rows and columns
cols=28;
rows=4;

//// bugs:
// - need to do this so that when you do double-sided, the dots don't
//   collide when you line up markers on front and back.
//   that means putting dots at + 0.5*dot2dot on opposite sides

// The thickness of the slate parts are important in 
// getting consistent dot placements; especially dot alignment
thickness = stylusDiameter;
pinheight=0.75;

//// rendering detail
// 15 degree features
$fa=15;
// preferred length of smallest feature
$fs=0.4;

//
// All parameters are derived after this
//

// The actual radius of the stylus is used a lot
stylusRadius=stylusDiameter/2;

// The distance from dot1 between lines is derived.
// The dot2dot and line extra are indirect control of this
line2line=(dots/2)*dot2dotV + lineExtra;

lockwidth=0.7;

marginCell=(cell2cell-dot2dotH)/2;
marginLine=(line2line-((dots/2)-1)*dot2dotV)/2;

//// This is the bottom part of the slate
// it features domed holes, and a punch to keep paper from moving,
// and a lock to ensure that it is aligned

translate([2,-cell2cell*cols/2,0])
union() {
    // These cones must line up when template is moved
    // Make them as short as possible to prevent breakage
    translate([stylusDiameter,cell2cell*(-1),thickness])
        cylinder(stylusDiameter*pinheight, stylusRadius, 0);
    translate([stylusDiameter,cell2cell*(cols+1),thickness])
        cylinder(stylusDiameter*pinheight, stylusRadius, 0);
    translate([line2line*rows+lineExtra,cell2cell*(-1),thickness])
        cylinder(stylusDiameter*pinheight, stylusRadius, 0);
    translate([line2line*rows+lineExtra,cell2cell*(cols+1),thickness])
        cylinder(stylusDiameter*pinheight, stylusRadius, 0);
    
    // Make the barrier that paper is placed against a trapezoid,
    // so that it does not get stuck on the template,
    // and template can flex around until it touches the backing.
    atx = (dot2dotV+lineExtra)/2;
    aty = cell2cell*(cols+2)-stylusDiameter;
    atz = 0;
    sx = line2line*rows;
    sy = stylusDiameter*lockwidth;
    sz = thickness*2;
    w = 10;
    difference() {
        union() {
            translate([atx,aty,atz])
                scale([sx, sy, sz])
                    cube(1);
        }
        union() {
            translate([atx-sx/4,aty+2,atz])
                rotate([20,0,0])
                    scale([sx*2, sy, 2*sz])
                        cube(1);
            translate([atx+sx,aty-w/2,atz+thickness])
                rotate([0,-20,0])
                scale([w,w,w])
                    cube(1);
            
        }
    }

    // This is the pits in the backing
    difference() {        
        translate([0,-2*cell2cell,0])
        scale([
          line2line*(rows)+dot2dotV+lineExtra,
          cell2cell*(cols+4),
          thickness
        ])
        cube(1);
        // This is so that double-sided notes don't collide
        translate([dot2dotV/2, -dot2dotH/2,0])
        for(c=[0:cols-1]) {
            for(r=[0:rows-1]) {
                for(cd=[0:1]) {
                    for(rd=[0:(dots/2)-1]) {
                        translate([
                            marginLine + r*line2line + rd*dot2dotV,
                            marginCell + c*cell2cell + cd*dot2dotH,
                            thickness
                        ])
                        scale([1,1,dome])
                        sphere(stylusRadius);
                    }
                }
            }
        }
    }
}

/// This is the template on top
// Notice that we inverted one axis to compensate for flipping the print over
// because the surface contacting is on top
translate([-4-line2line*rows-lineExtra,-cell2cell*cols/2,0])
difference() {
    // main backing
    union() {
        translate([0,-2*cell2cell,0])
        scale([
          line2line*(rows)+dot2dotV+lineExtra,
          cell2cell*(cols+4),
          thickness
        ])
        cube(1);
    }
    // drilled out items
    union() {
        // positioning pins
        translate([stylusDiameter,cell2cell*(-1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        translate([stylusDiameter,cell2cell*(cols+1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        translate([line2line*rows+lineExtra,cell2cell*(-1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        translate([line2line*rows+lineExtra,cell2cell*(cols+1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        
        // slot to align template
        translate([
            (dot2dotV+lineExtra)/2,
            cell2cell*(cols+2)-stylusDiameter,
            -thickness
        ])
        scale([line2line*(rows), stylusDiameter*lockwidth, 3*thickness])
                cube(1);
        
        // This offset is so that double-sided notes dont collide
        translate([dot2dotV/2,-dot2dotH/2,0])        
        for(c=[0:cols-1]) {
            for(r=[0:rows-1]) {
                for(cd=[0:1]) {                    
                    for(rd=[0:(dots/2)-1]) {
                        translate([
                            marginLine + r*line2line + rd*dot2dotV,
                            marginCell + c*cell2cell + cd*dot2dotH,
                            -thickness
                        ])
                        cylinder(
                            h=thickness*3,
                            r1=stylusRadius,
                            r2=stylusRadius
                        );
                    }
                }
                union() {
                    translate([
                        marginLine + r*line2line - stylusRadius,
                        marginCell + c*cell2cell - (1.6)*stylusDiameter/6,
                        -thickness
                    ])     
                    scale([
                        ((dots/2)-1)*dot2dotV + stylusDiameter,
                        dot2dotH + stylusRadius,
                        3*thickness
                    ])
                    cube(1);
                }
            }
        }        
    }
}

