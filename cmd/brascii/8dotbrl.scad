//// parameters of cell

// This is the millimeter unit. Multiply any lenghts times this
mm=1.0; 

// Set to this either 6 or 8, for the kind of braille
dots=8;

// standard spacing between dot centers
dot2dot=2.5*mm;

// extra spacing between lines (in addition to between dots
lineExtra=0.4*mm;

// spacing between cells at same dot
cell2cell=6.0*mm;

// the diameter of the stylus
// this is very important to get proper domes
stylusDiameter=2*mm;

// A value of 1.0 is a perfectly round dome.
// A perfectly round dome leaves rather large holes with paper rips
// A value of 0.75 gives a slightly flat ellipsoid pit
dome=0.75;

// parameters of slate in rows and columns
cols=28;
rows=4;

// The thickness of the slate parts are important in 
// getting consistent dot placements; especially dot alignment
thickness = stylusDiameter;

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

// When rendering more than 6-dots, make it easy to ignore the extra dots
// in the template, by making barriers around them larger
bigDots=6;

// The distance from dot1 between lines is derived.
// The dot2dot and line extra are indirect control of this
line2line=(dots/2)*dot2dot*mm + lineExtra;

fitmargin = 0.995;
lockwidth=0.85;

marginCell=(cell2cell-dot2dot)/2;
marginLine=(line2line-((dots/2)-1)*dot2dot)/2;

//// This is the bottom part of the slate
// it features domed holes, and a punch to keep paper from moving,
// and a lock to ensure that it is aligned

translate([1,-cell2cell*cols/2,0])
union() {
    // These cones must line up when template is moved
    // Make them as short as possible to prevent breakage
    translate([stylusDiameter,cell2cell*(-1),thickness])
        cylinder(stylusDiameter/2, stylusRadius, 0);
    translate([stylusDiameter,cell2cell*(cols+1),thickness])
        cylinder(stylusDiameter/2, stylusRadius, 0);
    translate([line2line*rows+lineExtra,cell2cell*(-1),thickness])
        cylinder(stylusDiameter/2, stylusRadius, 0);
    translate([line2line*rows+lineExtra,cell2cell*(cols+1),thickness])
        cylinder(stylusDiameter/2, stylusRadius, 0);
    
    translate([(dot2dot+lineExtra)/2,cell2cell*(cols+2)-stylusDiameter,0])
        scale([line2line*(rows)*fitmargin, stylusDiameter*lockwidth*fitmargin, thickness*2])
            cube(1);
    
    // invert this to be union to inspect the dome quality that should result.
    // there is an assumption that this dome matches your stylus tip.
    difference() {        
        translate([0,-2*cell2cell,0])
        scale([
          line2line*(rows)+dot2dot+lineExtra,
          cell2cell*(cols+4),
          thickness
        ])
        cube(1);
        for(c=[0:cols-1]) {
            for(r=[0:rows-1]) {
                for(cd=[0:1]) {
                    for(rd=[0:(dots/2)-1]) {
                        translate([
                            marginLine + r*line2line + rd*dot2dot,
                            marginCell + c*cell2cell + cd*dot2dot,
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
scale([-1,1,1])
translate([1,-cell2cell*cols/2,0])
difference() {
    union() {
        translate([0,-2*cell2cell,0])
        scale([
          line2line*(rows)+dot2dot+lineExtra,
          cell2cell*(cols+4),
          thickness
        ])
        cube(1);
    }
    union() {
        translate([stylusDiameter,cell2cell*(-1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        translate([stylusDiameter,cell2cell*(cols+1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        translate([line2line*rows+lineExtra,cell2cell*(-1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        translate([line2line*rows+lineExtra,cell2cell*(cols+1),-thickness])
            cylinder(3*thickness, stylusRadius, stylusRadius);
        
        translate([(dot2dot+lineExtra)/2,cell2cell*(cols+2)-stylusDiameter,-thickness])
            scale([line2line*(rows), stylusDiameter*lockwidth, 3*thickness])
                cube(1);
        
        for(c=[0:cols-1]) {
            for(r=[0:rows-1]) {
                for(cd=[0:1]) {                    
                    for(rd=[0:(dots/2)-1]) {
                        translate([
                            marginLine + r*line2line + rd*dot2dot,
                            marginCell + c*cell2cell + cd*dot2dot,
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
                        ((bigDots/2)-0.7)*dot2dot + stylusRadius,
                        dot2dot + stylusRadius,
                        3*thickness
                    ])
                    cube(1);
                    // hack! i have no idea
                    if(dots>6) {
                        translate([
                            marginLine + r*line2line - stylusDiameter/16,
                            marginCell + c*cell2cell - (0.5)*stylusDiameter/12,
                            -thickness
                        ])                    
                        scale([
                            ((dots/2)-0.5)*dot2dot - stylusDiameter/16,
                            dot2dot + stylusDiameter/24,
                            3*thickness
                        ])
                        cube(1);    
                    }                
                }
            }
        }        
    }
}

