let fs = require('fs')


// is this upper-case char
function IsUpperCaseChar(c) {
  if('A' <= c && c <= 'Z') {
    return true;
  }
  return false;
}


// c is upper-case
function ToLowerCaseChar(c) {
  return c.toLowerCase(c);
}

// is this lower-case char
function IsLowerCaseChar(c) {
  if('a' <= c && c <= 'z') {
    return true;
  }
  return false;
}

function IsAllUpperCaseWord(w) {
  if(w.length > 1) {
    allUc = true;
    for(i=0; i<w.length; i++) {
      if(IsLowerCaseChar(w[i])) {
        allUc = false;
      }
    }
    return allUc;
  }
  return false;
}

// n is a digit
function IsDigit(n) {
  if('0' <= n && n <= '9') {
    return true;
  }
  return false;
}

// n is a digit
function ToDigit(n) {
  return n;
}

// enums
const upperCase = 1;
const lowerCase = 2;
const digit = 3;

var brlWord = [];

// I am transliterating this:
// https://www.teachingvisuallyimpaired.com/uploads/1/4/1/2/14122361/ueb_braille_chart.pdf

  
var punctuation = [
  [",","1"],
  [".","4"],
  ["'","'"],
  [":","3"],
  ["!","6"],
  ["-","-"],
  ["?","8"],
  [";","2"],
  ["/","_/"],
  ["\\","_*"],
  ["(","\"<"],
  [")","\">"],
  ["[",".<"],
  ["]",".>"],
  ["{","_<"],
  ["}","_>"],
  ["<","`<"],
  [">","`>"],
  ["+","\"6"],
  ["-","\"-"],
  ["*","\"8"],
  ["/","\"/"],
  ["=","\"7"],
  ["?","@s"],
  ["%",".0"],
  ["#","_#"],
  ["&","@&"],
  ["@","@a"]
]; 

var aWordsigns = [
  ["but", "b"],
  ["can", "c"],
  ["do", "d"],
  ["every", "e"],
  ["from", "f"],
  ["go", "g"],
  ["have", "h"],
  ["just", "j"],
  ["knowledge", "k"],
  ["like", "l"],
  ["more", "m"],
  ["not", "n"],
  ["people", "p"],
  ["quite", "q"],
  ["rather", "r"],
  ["so", "s"],
  ["that", "t"],
  ["us", "u"],
  ["very", "v"],
  ["will", "w"],
  ["it", "x"],
  ["you", "y"],
  ["as", "z"]
];

var sWordsigns = [
  ["child","*"],
  ["shall","%"],
  ["this","?"],
  ["which",":"],
  ["out","|"],
  ["still","/"]
];

var lWordsigns = [
  ["be","2"],
  ["enough","5"],
  ["were","7"],
  ["his","8"],
  ["in","9"],
  ["was","0"]
];

var sContractions = [
  ["and","&"],
  ["for","="],
  ["of","("],
  ["the","!"],
  ["with",")"]
];

var sGroupsigns = [
  ["ch","*"],
  ["sh","%"],
  ["th","?"],
  ["wh",":"],
  ["ou","|"],
  ["st","/"],
  ["gh","<"],
  ["ed","?"],
  ["er","}"],
  ["ow","{"],
  ["ar",">"],
  ["ing","+"]
];

var lGroupsigns1 = [
  ["ea","'"],
  ["bb","2"],
  ["cc","3"],
  ["ff","6"],
  ["gg","7"]
];

var lGroupsigns2 = [
  ["be","2"],
  ["con","3"],
  ["dis","4"],
  ["en","5"],
  ["in","9"]
];

var lWordsigns = [
  ["be","2"],
  ["enough","5"],
  ["were","7"],
  ["his","8"],
  ["in","9"],
  ["was","0"]
];

var ilContractions = [
  ["day","\"d"],
  ["ever", "\"e"],
  ["father", "\"f"],
  ["here", "\"h"],
  ["know", "\"k"],
  ["lord", "\"l"],
  ["mother", "\"m"],
  ["name", "\"n"],
  ["one", "\"o"],
  ["part", "\"p"],
  ["question", "\"q"],
  ["right", "\"r"],
  ["some", "\"s"],
  ["time", "\"t"],
  ["under", "\"u"],
  ["work", "\"w"],
  ["young", "\"y"],
  ["there", "\"!"],
  ["character", "\"*"],
  ["through", "\"?"],
  ["where", "\":"],
  ["ought", "\"|"],
  ["upon", "`u"],
  ["word", "`w"],
  ["these", "`!"],
  ["those", "`?'"],
  ["whose", "`:"],
  ["cannot", "_c"],
  ["had", "_h"],
  ["many", "_m"],
  ["spirit", "_s"],
  ["world", "_w"],
  ["their", "_!"]
];

var flGroupsigns = [
  ["ound",".d"],
  ["ance",".e"],
  ["sion",".n"],
  ["less",".s"],
  ["ount",".t"],
  ["ence",";e"],
  ["ong",";g"],
  ["ful",";l"],
  ["tion",";n"],
  ["ness",";s"],
  ["ment",";t"],
  ["ity",";y"]
];

var sfWords = [
  ["about","ab"],
  ["above","abv"],
  ["according","ac"],
  ["across","acr"],
  ["after","af"],
  ["afternoon","afn"],
  ["afterward","afw"],
  ["again","ag"],
  ["against","ag/"],
  ["almost","alm"],
  ["already","alr"],
  ["also","al"],
  ["although","al?"],
  ["altogether","alt"],
  ["always","alw"],
  ["because","2c"],
  ["before","2f"],
  ["behind","2h"],
  ["below","2l"],
  ["beneath","2n"],
  ["beside","2s"],
  ["between","2t"], // conflict with actual word "bet" ? spell it out in braille to deconflict.
  ["beyond","2y"],
  ["blind","bl"],
  ["braille","brl"],
  ["children","*n"],
  ["conceive","3cv"],
  ["conceiving","3cvg"],
  ["could","cd"],
  ["deceive","dcv"],
  ["deceiving","dcvg"],
  ["declare","dcl"],
  ["declaring","dclg"],
  ["either","ei"],
  ["first","f*"],
  ["friend","fr"],
  ["good","gd"],
  ["great","grt"],
  ["herself","h}f"],
  ["him","hm"],
  ["himself","hmf"],
  ["immediate","imm"],
  ["its","xs"],
  ["itself","xf"],
  ["letter","lt"],
  ["little","ll"],
  ["much","m*"],
  ["must","m/"],
  ["myself","myf"],
  ["necessary","nec"],
  ["neither","nei"],
  ["oneself","\"of"],
  ["ourselves","|rvs"],
  ["paid","pd"],
  ["perceive","p}cv"],
  ["perceiving","p}cvg"],
  ["perhaps","p}h"],
  ["quick","qk"],
  ["receive","rcv"],
  ["receiving","rcvg"],
  ["rejoice","rjc"],
  ["rejoicing","rjcg"],
  ["said","sd"],
  ["should","/d"],
  ["such","s*"],
  ["themselves","!mvs"],
  ["thyself","?yf"],
  ["today","td"],
  ["together","tgr"],
  ["tomorrow","tm"],
  ["tonight","tn"],
  ["would","wd"],
  ["your","yr"],
  ["yourself","yrf"],
  ["yourselves","yrsv"] 
];

function findFirstFwd(w, table) {
  for(var i=0; i < table.length; i++) {
    var t = table[i];
    var k = t[0];
    var v = t[1];
    if(w == k) {
      return v;
    } 
  }
  return null;
}

function substInitialFwd(w) {
  for(var i=0; i < ilContractions.length; i++) {
    var t = ilContractions[i];
    var k = t[0];
    var v = t[1];
    if(w.startsWith(k)) {
      return [v, w.substring(k.length)];
    } 
  }
  return null;
}

function substFinalFwd(w) {
  for(var i=0; i < flGroupsigns.length; i++) {
    var t = flGroupsigns[i];
    var k = t[0];
    var v = t[1];
    if(w.endsWith(k)) {
      var pre = w.substring(0, w.length - k.length);
      return [pre,v];
    } 
  }
  return null;
}

function compressMiddle(w) {
  tables = [
    sContractions,
    sGroupsigns,
    lGroupsigns1,
    lGroupsigns2
  ];
  for(var t = 0; t < tables.length; t++) {
    var table = tables[t];
    while(true) {
      // do all subst for this table until none can be made
      var oldw = w;
      for(var i=0; i < table.length; i++) {
        var key = table[i][0];
        var val = table[i][1];
        var w = w.replace(key,val);
      }
      if(w == oldw) {
        break;
      }
    }
  }
  return w;
}

// I think that all of ueb only compresses single words.
function compressWord(w) {
  // per-word caps and number handling
  var prefix = "";
  var start = "";
  var middle = w;
  var end = "";

  if(IsAllUpperCaseWord(middle)) {
    prefix = ",,";
    middle = middle.toLowerCase();
  } else {
    if(IsUpperCaseChar(middle[0])) {
      prefix = ",";
      middle = middle.toLowerCase()
    } else {
      if(IsDigit(middle[0])) {
        prefix = "#";
        // TODO
      }
    }
  }

  // whole word substitutions
  var v = findFirstFwd(middle,sfWords);
  if(v != null) {
    return [prefix, v];
  }

  // single-char substitutions
  var v = findFirstFwd(middle,aWordsigns);
  if(v != null) {
    return [prefix, v];
  }

  var v = findFirstFwd(middle,sWordsigns);
  if(v != null) {
    return [prefix, v];
  }

  var v = findFirstFwd(middle,lWordsigns);
  if(v != null) {
    return [prefix, v];
  }
 
  // Split out the start 
  if(middle.length > 0) {
    var sv = substInitialFwd(middle);
    if(sv != null) {
      start = sv[0];
      middle = sv[1]; 
    }
  }

  // Split out ending
  if(middle.length > 0) {
    var ev = substFinalFwd(middle);
    if(ev != null) {
      middle = ev[0];
      end = ev[1];  
    }
  }

  middle = compressMiddle(middle);

  return [prefix, start+middle+end];
}

function compressWord2(w) {
  wd = compressWord(w);
  return (wd[0] + wd[1]);
}


/*
sentence = [
 "prompt",
 "engineering"
];

for(i=0; i<sentence.length; i++) {
  process.stdout.write(sentence[i] + " ");
}
process.stdout.write("\n");
for(i=0; i<sentence.length; i++) {
  process.stdout.write(compressWord(sentence[i]) + " ");
}
*/

sentence = [
  "Braille","is","a","crazy","language","because","engineering","the","compression","is","farking", "painful"
];

for(var j=0; j<sentence.length; j++) {
  process.stdout.write(compressWord2(sentence[j]) + " ");
}

