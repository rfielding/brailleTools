GoatRope
=========

[goat.jpg](goat.jpg)

In building text editors, and stream editors; one of the most fundamental tasks
is to build an object that handles easy inserts and deletes into a stream.
One approach is to make this object looks like a file; specifically, to look
like a character device driver.  The advantage of this approach is that we
have the os.File like interface as a starting point.

> The Unix file interface does not have simple methods to handle inserts and deletes.

From there, it is not difficult to navigate (ie: seek) the file with an
awareness of carriage returns.  Once line-breaks are handled, cursor
directions such as up/down/left/right make sense, as does reading in a line-oriented way.

The GoatRope is:

- Based on a PieceTable, which was one of the things that got early Microsoft Word off the ground.
- A completely separate package from the Braille text editor which is driving it as a specific need.

If you need to be able to insert and delete into large files efficiently, or edit file streams;
you may find this package useful.  You can use the PieceTable on its own, as it is completely separate
from file IO, and tested on its own.  You make GoatRopes backed by files or memory.

A Braille editor is an extreme example of a line-oriented editor.  It makes sense to have the GoatRope
be a file enhanced with navigation, search, and index functions; where the editor that wraps around it
concerns itself completely with ANSI escape sequences and character generation (ie: building characters
out of dots in Braille, should not creep into the GoatRope, but navigation, search, and indexing probably should).
