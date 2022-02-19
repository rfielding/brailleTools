## Braille Binary

> The 6-dot standard is 0x20 to 0x5F with dot7 masked off.
> Bottom half of standard is used for control codes 0x00 to 0x19.

|         |second _0|second _1|second _2|second _3|second _4|second _5|second _6|second _7|second _8|second _9|second _A|second _B|second _C|second _D|second _E|second _F|
|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|---------|
|       0_| ⡀ NUL   | ⡮ SOH   | ⡐ STX   | ⡼ ETX   | ⡫ EOT   | ⡩ ENQ   | ⡯ ACK   | ⡄ BEL   | ⡷  BS   | ⡾ TAB   | ⡡  LF   | ⡬  VT   | ⡠  FF   | ⡤  CR   | ⡨  SO   | ⡌  SI   |
|       1_| ⡴ DLE   | ⡂ DC1   | ⡆ DC2   | ⡒ DC3   | ⡲ DC4   | ⡢ NAK   | ⡖ SYN   | ⡶ ETB   | ⡦ CAN   | ⡔  EM   | ⡱ SUB   | ⡰ ESC   | ⡣  FS   | ⡿  GS   | ⡜  RS   | ⡹  US   |
|       2_| ⠀ SPC   | ⠮   !   | ⠐   "   | ⠼   #   | ⠫   $   | ⠩   %   | ⠯   &   | ⠄   '   | ⠷   (   | ⠾   )   | ⠡   *   | ⠬   +   | ⠠   ,   | ⠤   -   | ⠨   .   | ⠌   /   |
|       3_| ⠴   0   | ⠂   1   | ⠆   2   | ⠒   3   | ⠲   4   | ⠢   5   | ⠖   6   | ⠶   7   | ⠦   8   | ⠔   9   | ⠱   :   | ⠰   ;   | ⠣   <   | ⠿   =   | ⠜   >   | ⠹   ?   |
|       4_| ⠈   @   | ⡁   A   | ⡃   B   | ⡉   C   | ⡙   D   | ⡑   E   | ⡋   F   | ⡛   G   | ⡓   H   | ⡊   I   | ⡚   J   | ⡅   K   | ⡇   L   | ⡍   M   | ⡝   N   | ⡕   O   |
|       5_| ⡏   P   | ⡟   Q   | ⡗   R   | ⡎   S   | ⡞   T   | ⡥   U   | ⡧   V   | ⡺   W   | ⡭   X   | ⡽   Y   | ⡵   Z   | ⠪   [   | ⠳   \   | ⠻   ]   | ⠘   ^   | ⠸   _   |
|       6_| ⡈   `   | ⠁   a   | ⠃   b   | ⠉   c   | ⠙   d   | ⠑   e   | ⠋   f   | ⠛   g   | ⠓   h   | ⠊   i   | ⠚   j   | ⠅   k   | ⠇   l   | ⠍   m   | ⠝   n   | ⠕   o   |
|       7_| ⠏   p   | ⠟   q   | ⠗   r   | ⠎   s   | ⠞   t   | ⠥   u   | ⠧   v   | ⠺   w   | ⠭   x   | ⠽   y   | ⠵   z   | ⡪   {   | ⡳  \|   | ⡻   }   | ⡘   ~   | ⡸ DEL   |
|       8_| ⣀ UNK   | ⣮ UNK   | ⣐ UNK   | ⣼ UNK   | ⣫ UNK   | ⣩ UNK   | ⣯ UNK   | ⣄ UNK   | ⣷ UNK   | ⣾ UNK   | ⣡ UNK   | ⣬ UNK   | ⣠ UNK   | ⣤ UNK   | ⣨ UNK   | ⣌ UNK   |
|       9_| ⣴ UNK   | ⣂ UNK   | ⣆ UNK   | ⣒ UNK   | ⣲ UNK   | ⣢ UNK   | ⣖ UNK   | ⣶ UNK   | ⣦ UNK   | ⣔ UNK   | ⣱ UNK   | ⣰ UNK   | ⣣ UNK   | ⣿ UNK   | ⣜ UNK   | ⣹ UNK   |
|       A_| ⢀ UNK   | ⢮   ¡   | ⢐   ¢   | ⢼   £   | ⢫   ¤   | ⢩   ¥   | ⢯   ¦   | ⢄   §   | ⢷   ¨   | ⢾   ©   | ⢡   ª   | ⢬   «   | ⢠   ¬   | ⢤   ­   | ⢨   ®   | ⢌   ¯   |
|       B_| ⢴   °   | ⢂   ±   | ⢆   ²   | ⢒   ³   | ⢲   ´   | ⢢   µ   | ⢖   ¶   | ⢶   ·   | ⢦   ¸   | ⢔   ¹   | ⢱   º   | ⢰   »   | ⢣   ¼   | ⢿   ½   | ⢜   ¾   | ⢹   ¿   |
|       C_| ⢈   À   | ⣁   Á   | ⣃   Â   | ⣉   Ã   | ⣙   Ä   | ⣑   Å   | ⣋   Æ   | ⣛   Ç   | ⣓   È   | ⣊   É   | ⣚   Ê   | ⣅   Ë   | ⣇   Ì   | ⣍   Í   | ⣝   Î   | ⣕   Ï   |
|       D_| ⣏   Ð   | ⣟   Ñ   | ⣗   Ò   | ⣎   Ó   | ⣞   Ô   | ⣥   Õ   | ⣧   Ö   | ⣺   ×   | ⣭   Ø   | ⣽   Ù   | ⣵   Ú   | ⢪   Û   | ⢳   Ü   | ⢻   Ý   | ⢘   Þ   | ⢸   ß   |
|       E_| ⣈   à   | ⢁   á   | ⢃   â   | ⢉   ã   | ⢙   ä   | ⢑   å   | ⢋   æ   | ⢛   ç   | ⢓   è   | ⢊   é   | ⢚   ê   | ⢅   ë   | ⢇   ì   | ⢍   í   | ⢝   î   | ⢕   ï   |
|       F_| ⢏   ð   | ⢟   ñ   | ⢗   ò   | ⢎   ó   | ⢞   ô   | ⢥   õ   | ⢧   ö   | ⢺   ÷   | ⢭   ø   | ⢽   ù   | ⢵   ú   | ⣪   û   | ⣳   ü   | ⣻   ý   | ⣘   þ   | ⣸   ÿ   |
|
