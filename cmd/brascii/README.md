## Braille Binary

> The 6-dot standard is 0x20 to 0x5F with dot7 masked off.
> Bottom half of standard is used for control codes 0x00 to 0x19.

|   |     _0|     _1|     _2|     _3|     _4|     _5|     _6|     _7|     _8|     _9|     _A|     _B|     _C|     _D|     _E|     _F|
|---|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| 0_| ⡀ <br> NUL | ⡮ <br> SOH | ⡐ <br> STX | ⡼ <br> ETX | ⡫ <br> EOT | ⡩ <br> ENQ | ⡯ <br> ACK | ⡄ <br> BEL | ⡷ <br>  BS | ⡾ <br> TAB | ⡡ <br>  LF | ⡬ <br>  VT | ⡠ <br>  FF | ⡤ <br>  CR | ⡨ <br>  SO | ⡌ <br>  SI |
| 1_| ⡴ <br> DLE | ⡂ <br> DC1 | ⡆ <br> DC2 | ⡒ <br> DC3 | ⡲ <br> DC4 | ⡢ <br> NAK | ⡖ <br> SYN | ⡶ <br> ETB | ⡦ <br> CAN | ⡔ <br>  EM | ⡱ <br> SUB | ⡰ <br> ESC | ⡣ <br>  FS | ⡿ <br>  GS | ⡜ <br>  RS | ⡹ <br>  US |
| 2_| ⠀ <br> SPC | ⠮ <br>   ! | ⠐ <br>   " | ⠼ <br>   # | ⠫ <br>   $ | ⠩ <br>   % | ⠯ <br>   & | ⠄ <br>   ' | ⠷ <br>   ( | ⠾ <br>   ) | ⠡ <br>   * | ⠬ <br>   + | ⠠ <br>   , | ⠤ <br>   - | ⠨ <br>   . | ⠌ <br>   / |
| 3_| ⠴ <br>   0 | ⠂ <br>   1 | ⠆ <br>   2 | ⠒ <br>   3 | ⠲ <br>   4 | ⠢ <br>   5 | ⠖ <br>   6 | ⠶ <br>   7 | ⠦ <br>   8 | ⠔ <br>   9 | ⠱ <br>   : | ⠰ <br>   ; | ⠣ <br>   < | ⠿ <br>   = | ⠜ <br>   > | ⠹ <br>   ? |
| 4_| ⠈ <br>   @ | ⡁ <br>   A | ⡃ <br>   B | ⡉ <br>   C | ⡙ <br>   D | ⡑ <br>   E | ⡋ <br>   F | ⡛ <br>   G | ⡓ <br>   H | ⡊ <br>   I | ⡚ <br>   J | ⡅ <br>   K | ⡇ <br>   L | ⡍ <br>   M | ⡝ <br>   N | ⡕ <br>   O |
| 5_| ⡏ <br>   P | ⡟ <br>   Q | ⡗ <br>   R | ⡎ <br>   S | ⡞ <br>   T | ⡥ <br>   U | ⡧ <br>   V | ⡺ <br>   W | ⡭ <br>   X | ⡽ <br>   Y | ⡵ <br>   Z | ⠪ <br>   [ | ⠳ <br>   \ | ⠻ <br>   ] | ⠘ <br>   ^ | ⠸ <br>   _ |
| 6_| ⡈ <br>   ` | ⠁ <br>   a | ⠃ <br>   b | ⠉ <br>   c | ⠙ <br>   d | ⠑ <br>   e | ⠋ <br>   f | ⠛ <br>   g | ⠓ <br>   h | ⠊ <br>   i | ⠚ <br>   j | ⠅ <br>   k | ⠇ <br>   l | ⠍ <br>   m | ⠝ <br>   n | ⠕ <br>   o |
| 7_| ⠏ <br>   p | ⠟ <br>   q | ⠗ <br>   r | ⠎ <br>   s | ⠞ <br>   t | ⠥ <br>   u | ⠧ <br>   v | ⠺ <br>   w | ⠭ <br>   x | ⠽ <br>   y | ⠵ <br>   z | ⡪ <br>   { | ⡳ <br>  \| | ⡻ <br>   } | ⡘ <br>   ~ | ⡸ <br> DEL |
| 8_| ⣀ <br> UNK | ⣮ <br> UNK | ⣐ <br> UNK | ⣼ <br> UNK | ⣫ <br> UNK | ⣩ <br> UNK | ⣯ <br> UNK | ⣄ <br> UNK | ⣷ <br> UNK | ⣾ <br> UNK | ⣡ <br> UNK | ⣬ <br> UNK | ⣠ <br> UNK | ⣤ <br> UNK | ⣨ <br> UNK | ⣌ <br> UNK |
| 9_| ⣴ <br> UNK | ⣂ <br> UNK | ⣆ <br> UNK | ⣒ <br> UNK | ⣲ <br> UNK | ⣢ <br> UNK | ⣖ <br> UNK | ⣶ <br> UNK | ⣦ <br> UNK | ⣔ <br> UNK | ⣱ <br> UNK | ⣰ <br> UNK | ⣣ <br> UNK | ⣿ <br> UNK | ⣜ <br> UNK | ⣹ <br> UNK |
| A_| ⢀ <br> UNK | ⢮ <br>   ¡ | ⢐ <br>   ¢ | ⢼ <br>   £ | ⢫ <br>   ¤ | ⢩ <br>   ¥ | ⢯ <br>   ¦ | ⢄ <br>   § | ⢷ <br>   ¨ | ⢾ <br>   © | ⢡ <br>   ª | ⢬ <br>   « | ⢠ <br>   ¬ | ⢤ <br>   ­ | ⢨ <br>   ® | ⢌ <br>   ¯ |
| B_| ⢴ <br>   ° | ⢂ <br>   ± | ⢆ <br>   ² | ⢒ <br>   ³ | ⢲ <br>   ´ | ⢢ <br>   µ | ⢖ <br>   ¶ | ⢶ <br>   · | ⢦ <br>   ¸ | ⢔ <br>   ¹ | ⢱ <br>   º | ⢰ <br>   » | ⢣ <br>   ¼ | ⢿ <br>   ½ | ⢜ <br>   ¾ | ⢹ <br>   ¿ |
| C_| ⢈ <br>   À | ⣁ <br>   Á | ⣃ <br>   Â | ⣉ <br>   Ã | ⣙ <br>   Ä | ⣑ <br>   Å | ⣋ <br>   Æ | ⣛ <br>   Ç | ⣓ <br>   È | ⣊ <br>   É | ⣚ <br>   Ê | ⣅ <br>   Ë | ⣇ <br>   Ì | ⣍ <br>   Í | ⣝ <br>   Î | ⣕ <br>   Ï |
| D_| ⣏ <br>   Ð | ⣟ <br>   Ñ | ⣗ <br>   Ò | ⣎ <br>   Ó | ⣞ <br>   Ô | ⣥ <br>   Õ | ⣧ <br>   Ö | ⣺ <br>   × | ⣭ <br>   Ø | ⣽ <br>   Ù | ⣵ <br>   Ú | ⢪ <br>   Û | ⢳ <br>   Ü | ⢻ <br>   Ý | ⢘ <br>   Þ | ⢸ <br>   ß |
| E_| ⣈ <br>   à | ⢁ <br>   á | ⢃ <br>   â | ⢉ <br>   ã | ⢙ <br>   ä | ⢑ <br>   å | ⢋ <br>   æ | ⢛ <br>   ç | ⢓ <br>   è | ⢊ <br>   é | ⢚ <br>   ê | ⢅ <br>   ë | ⢇ <br>   ì | ⢍ <br>   í | ⢝ <br>   î | ⢕ <br>   ï |
| F_| ⢏ <br>   ð | ⢟ <br>   ñ | ⢗ <br>   ò | ⢎ <br>   ó | ⢞ <br>   ô | ⢥ <br>   õ | ⢧ <br>   ö | ⢺ <br>   ÷ | ⢭ <br>   ø | ⢽ <br>   ù | ⢵ <br>   ú | ⣪ <br>   û | ⣳ <br>   ü | ⣻ <br>   ý | ⣘ <br>   þ | ⣸ <br>   ÿ |
|
