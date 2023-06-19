let fs = require('fs')
let readline = require('readline')


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

function IsChar(c) {
  return IsLowerCaseChar(c) || IsUpperCaseChar(c);
}

function IsAllUpperCaseWord(w) {
  if(w.length > 1) {
    allUc = true;
    for(var i=0; i<w.length; i++) {
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

// Computer Braille is a permutation of locations 0-255, where
// in practice, 1-127 are used for printable ASCII as a braille font.
const brailleAsciiPattern = [
  0b00000000, 0b00101110, 0b00010000, 0b00111100, 0b00101011, 0b00101001, 0b00101111, 0b00000100, 0b00110111, 0b00111110, 0b00100001, 0b00101100, 0b00100000, 0b00100100, 0b00101000, 0b00001100,
  0b00110100, 0b00000010, 0b00000110, 0b00010010, 0b00110010, 0b00100010, 0b00010110, 0b00110110, 0b00100110, 0b00010100, 0b00110001, 0b00110000, 0b00100011, 0b00111111, 0b00011100, 0b00111001,
  0b00001000, 0b00000001, 0b00000011, 0b00001001, 0b00011001, 0b00010001, 0b00001011, 0b00011011, 0b00010011, 0b00001010, 0b00011010, 0b00000101, 0b00000111, 0b00001101, 0b00011101, 0b00010101,
  0b00001111, 0b00011111, 0b00010111, 0b00001110, 0b00011110, 0b00100101, 0b00100111, 0b00111010, 0b00101101, 0b00111101, 0b00110101, 0b00101010, 0b00110011, 0b00111011, 0b00011000, 0b00111000,
];

// When looking at 0x20 through 0x5F as 6-dot, mask off dot 7 first,
// as there are only upper-case letters in braille ascii
const braillePerm = new Array(256);  
const asciiPerm = new Array(256);

function brailleInit() {
  const present = new Array(256);

  // Copy in the standard braille ascii patern
  for(var i = 0; i < 64; i++) {
    braillePerm[0x20+i] = brailleAsciiPattern[i]
  }
  // Flip the case of the alphabet
  for(var i = 0x40; i <= 0x60; i++) {
    braillePerm[i] = braillePerm[i] ^ 0x40
  }
  // Copy lower half of standard to cover control codes
  for(var i = 0; i < 32; i++) {
    braillePerm[i] = (braillePerm[i+0x20]) ^ 0x40
  }
  // Copy upper half of standard to cover upper case
  for(var i = 0; i < 32; i++) {
    braillePerm[0x60+i] = braillePerm[0x40+i] ^ 0x40
  }
  // Swap 124 and 127 the underscore and delete,
  // a strange exception logically, but I see it in real terminals
  var braillePermTmp = braillePerm[0x7F]
  braillePerm[0x7F] = braillePerm[0x5F]
  braillePerm[0x5F] = braillePermTmp

  // Duplicated it all in high bits
  for(var i = 0; i < 128; i++) {
    braillePerm[0x80+i] = braillePerm[i] ^ 0x80
  }
  // Reverse mapping
  for(var i = 0; i < 256; i++) {
    asciiPerm[braillePerm[i]] = i
    present[i]=1;
  }
  // Panic if codes are missing or duplicated
  for(var i = 0; i < 256; i++) {
    if(present[i] != 1) {
      console.log("braille table inconsistency at %d", i);
    }
  }
}


// I am transliterating this:
// https://www.teachingvisuallyimpaired.com/uploads/1/4/1/2/14122361/ueb_braille_chart.pdf

const digitMap = [
  ["1","a"],
  ["2","b"],
  ["3","c"],
  ["4","d"],
  ["5","e"],
  ["6","f"],
  ["7","g"],
  ["8","h"],
  ["9","i"],
  ["0","j"],
];
  
const punctuation = [
  ["_",",_"],
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
  ["%",".0"],
  ["#","_#"],
  ["&","`&"],
  ["@","`a"],
  ["$","`s"]
]; 

const aWordsigns = [
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

const sWordsigns = [
  ["child","*"],
  ["shall","%"],
  ["this","?"],
  ["which",":"],
  ["out","|"],
  ["still","/"]
];

const lWordsigns = [
  ["be","2"],
  ["enough","5"],
  ["were","7"],
  ["his","8"],
  ["in","9"],
  ["was","0"]
];

const sContractions = [
  ["and","&"],
  ["for","="],
  ["of","("],
  ["the","!"],
  ["with",")"]
];

const sGroupsigns = [
  ["ch","*"],
  ["sh","%"],
  ["th","?"],
  ["wh",":"],
  ["ou","|"],
  ["st","/"],
  ["gh","<"],
  ["ed","$"],
  ["er","}"],
  ["ow","{"],
  ["ar",">"],
  ["ing","+"]
];

const lGroupsigns1 = [
  ["ea","1"],
  ["bb","2"],
  ["cc","3"],
  ["ff","6"],
  ["gg","7"]
];

const lGroupsigns2 = [
  ["be","2"],
  ["con","3"],
  ["dis","4"],
  ["en","5"],
  ["in","9"]
];

const ilContractions = [
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

const flGroupsigns = [
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

const sfWords = [
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
    ilContractions,
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
  var wd = compressWord(w);
  return (wd[0] + wd[1]);
}


function translateString(w) {
  var out = [];
  var wordChars = [];
  var digits = [];
  var isQuoting = false;

  var flush = function(j) {
    if(wordChars.length > 0) {
      out.push(compressWord2(wordChars.join("")));
      wordChars = [];
    } else if(digits.length > 0) {
      var v = "#";
      if(w.length > 2) {
        j--;
        while(j>0 && IsDigit(w[j])) {
          j--;
        } 
        if(j > 1 && (j+1)<w.length && w[j] == "." && IsDigit(w[j+1]) && IsDigit(w[j-1])) {
          v = "";
        }
      }
      for(var i=0; i<digits.length; i++) {
        v = v + findFirstFwd(digits[i],digitMap);
      }
      out.push(v);
      digits = [];
    }
  };

  if(w.length > 0) {
    var wasReadingToken = IsChar(w[0]);
    var wasReadingNumber = IsDigit(w[0]);
    for(var i=0; i<w.length; i++) {
      var c = w[i];
      var isReadingToken  = IsChar(c); 
      var isReadingNumber = IsDigit(c);
      if(isReadingToken) {
        if(wasReadingNumber) {
          flush(i);
          if(IsLowerCaseChar(c)) {
            out.push(";");
          }
        }
        wordChars.push(c);
      } else if(isReadingNumber) {
        if(wasReadingToken) {
          flush(i);
        }
        digits.push(c);
      } else {
        flush(i);
        var f = findFirstFwd(c, punctuation);
        if(f != null) {
          out.push(f);
        } else if(c == " " || c == "\r" || c == "\n" || c == "\t") {
            out.push(c);
        } else if(c == "\"") {
          if(isQuoting) {
            out.push("0");
            isQuoting = false;
          } else {
            out.push("8");
            isQuoting = true;
          }
        } else {
          // Is there anything to ignore here?
          out.push(c);
        }
      }
      wasReadingToken = isReadingToken;
      wasReadingNumber = isReadingNumber;
    }
    flush(i); 
  }
  return out.join("");
}

var rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

//
// Do main work here
//

// Init and process args
brailleInit();
var asText = false;
var decode = false;
for(var i=2; i < process.argv.length; i++) {
  if(process.argv[i] == "--asText") {
    asText = true;
  }
  if(process.argv[i] == "--decode") {
    decode = true;
  }
}

rl.on('line', function (line) {
  if(decode) {
    console.log("TODO: implement braille decode");
  } else {
    var brl = translateString(line);
    var out = [];
    for(var i=0; i<brl.length; i++) {
      if(asText) {
        out.push(brl[i]);
      } else {
        out.push(String.fromCharCode(braillePerm[brl.charCodeAt(i)]+0x2800));
      }
    }
    console.log(out.join(""));
  }
});



