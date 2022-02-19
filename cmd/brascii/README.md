## Braille Binary

> The 6-dot standard is 0x20 to 0x5F with dot7 masked off.
> Bottom half of standard is used for control codes 0x00 to 0x19.

|         |       _0|       _1|       _2|       _3|       _4|       _5|       _6|       _7|       _8|       _9|       _A|       _B|       _C|       _D|       _E|       _F|
|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|
<nobr>|       0_| ⡀ NUL   | ⡮ SOH   | ⡐ STX   | ⡼ ETX   | ⡫ EOT   | ⡩ ENQ   | ⡯ ACK   | ⡄ BEL   | ⡷  BS   | ⡾ TAB   | ⡡  LF   | ⡬  VT   | ⡠  FF   | ⡤  CR   | ⡨  SO   | ⡌  SI   </nobr>|
<nobr>|       1_| ⡴ DLE   | ⡂ DC1   | ⡆ DC2   | ⡒ DC3   | ⡲ DC4   | ⡢ NAK   | ⡖ SYN   | ⡶ ETB   | ⡦ CAN   | ⡔  EM   | ⡱ SUB   | ⡰ ESC   | ⡣  FS   | ⡿  GS   | ⡜  RS   | ⡹  US   </nobr>|
<nobr>|       2_| ⠀ SPC   | ⠮   !   | ⠐   "   | ⠼   #   | ⠫   $   | ⠩   %   | ⠯   &   | ⠄   '   | ⠷   (   | ⠾   )   | ⠡   *   | ⠬   +   | ⠠   ,   | ⠤   -   | ⠨   .   | ⠌   /   </nobr>|
<nobr>|       3_| ⠴   0   | ⠂   1   | ⠆   2   | ⠒   3   | ⠲   4   | ⠢   5   | ⠖   6   | ⠶   7   | ⠦   8   | ⠔   9   | ⠱   :   | ⠰   ;   | ⠣   <   | ⠿   =   | ⠜   >   | ⠹   ?   </nobr>|
<nobr>|       4_| ⠈   @   | ⡁   A   | ⡃   B   | ⡉   C   | ⡙   D   | ⡑   E   | ⡋   F   | ⡛   G   | ⡓   H   | ⡊   I   | ⡚   J   | ⡅   K   | ⡇   L   | ⡍   M   | ⡝   N   | ⡕   O   </nobr>|
<nobr>|       5_| ⡏   P   | ⡟   Q   | ⡗   R   | ⡎   S   | ⡞   T   | ⡥   U   | ⡧   V   | ⡺   W   | ⡭   X   | ⡽   Y   | ⡵   Z   | ⠪   [   | ⠳   \   | ⠻   ]   | ⠘   ^   | ⠸   _   </nobr>|
<nobr>|       6_| ⡈   `   | ⠁   a   | ⠃   b   | ⠉   c   | ⠙   d   | ⠑   e   | ⠋   f   | ⠛   g   | ⠓   h   | ⠊   i   | ⠚   j   | ⠅   k   | ⠇   l   | ⠍   m   | ⠝   n   | ⠕   o   </nobr>|
<nobr>|       7_| ⠏   p   | ⠟   q   | ⠗   r   | ⠎   s   | ⠞   t   | ⠥   u   | ⠧   v   | ⠺   w   | ⠭   x   | ⠽   y   | ⠵   z   | ⡪   {   | ⡳  \|   | ⡻   }   | ⡘   ~   | ⡸ DEL   </nobr>|
<nobr>|       8_| ⣀ UNK   | ⣮ UNK   | ⣐ UNK   | ⣼ UNK   | ⣫ UNK   | ⣩ UNK   | ⣯ UNK   | ⣄ UNK   | ⣷ UNK   | ⣾ UNK   | ⣡ UNK   | ⣬ UNK   | ⣠ UNK   | ⣤ UNK   | ⣨ UNK   | ⣌ UNK   </nobr>|
<nobr>|       9_| ⣴ UNK   | ⣂ UNK   | ⣆ UNK   | ⣒ UNK   | ⣲ UNK   | ⣢ UNK   | ⣖ UNK   | ⣶ UNK   | ⣦ UNK   | ⣔ UNK   | ⣱ UNK   | ⣰ UNK   | ⣣ UNK   | ⣿ UNK   | ⣜ UNK   | ⣹ UNK   </nobr>|
<nobr>|       A_| ⢀ UNK   | ⢮   ¡   | ⢐   ¢   | ⢼   £   | ⢫   ¤   | ⢩   ¥   | ⢯   ¦   | ⢄   §   | ⢷   ¨   | ⢾   ©   | ⢡   ª   | ⢬   «   | ⢠   ¬   | ⢤   ­   | ⢨   ®   | ⢌   ¯   </nobr>|
<nobr>|       B_| ⢴   °   | ⢂   ±   | ⢆   ²   | ⢒   ³   | ⢲   ´   | ⢢   µ   | ⢖   ¶   | ⢶   ·   | ⢦   ¸   | ⢔   ¹   | ⢱   º   | ⢰   »   | ⢣   ¼   | ⢿   ½   | ⢜   ¾   | ⢹   ¿   </nobr>|
<nobr>|       C_| ⢈   À   | ⣁   Á   | ⣃   Â   | ⣉   Ã   | ⣙   Ä   | ⣑   Å   | ⣋   Æ   | ⣛   Ç   | ⣓   È   | ⣊   É   | ⣚   Ê   | ⣅   Ë   | ⣇   Ì   | ⣍   Í   | ⣝   Î   | ⣕   Ï   </nobr>|
<nobr>|       D_| ⣏   Ð   | ⣟   Ñ   | ⣗   Ò   | ⣎   Ó   | ⣞   Ô   | ⣥   Õ   | ⣧   Ö   | ⣺   ×   | ⣭   Ø   | ⣽   Ù   | ⣵   Ú   | ⢪   Û   | ⢳   Ü   | ⢻   Ý   | ⢘   Þ   | ⢸   ß   </nobr>|
<nobr>|       E_| ⣈   à   | ⢁   á   | ⢃   â   | ⢉   ã   | ⢙   ä   | ⢑   å   | ⢋   æ   | ⢛   ç   | ⢓   è   | ⢊   é   | ⢚   ê   | ⢅   ë   | ⢇   ì   | ⢍   í   | ⢝   î   | ⢕   ï   </nobr>|
<nobr>|       F_| ⢏   ð   | ⢟   ñ   | ⢗   ò   | ⢎   ó   | ⢞   ô   | ⢥   õ   | ⢧   ö   | ⢺   ÷   | ⢭   ø   | ⢽   ù   | ⢵   ú   | ⣪   û   | ⣳   ü   | ⣻   ý   | ⣘   þ   | ⣸   ÿ   </nobr>|
|
