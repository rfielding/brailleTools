GoatRope
=========

![goat.jpg](goat.jpg)

In building text editors, and stream editors; one of the most fundamental tasks
is to build an object that handles easy inserts and deletes into a stream.
One approach is to make this object looks like a file; specifically, to look
like a character device driver.  The advantage of this approach is that we
have the os.File like interface as a starting point.

> The Unix file interface does not have simple methods to handle inserts and deletes.

From there, it is not difficult to navigate (ie: seek) the file with an
awareness of carriage returns.  Once line-breaks are handled, cursor
directions such as up/down/left/right make sense, as does reading in a line-oriented way.

There seems to be a dearth of readily available simple starting points for 
editing files as a stream, so I will keep GoatRope separate for appropriate use
for anybody that needs such a data structure.  The main problem with:

```go
type File interface {
	io.Reader
	io.Seeker
	io.Closer
	io.Writer
	Stat() (os.FileInfo, error)
}
```

Is that it's not well defined how to insert into a large file, because you can't just shift
every character to the right one spot efficiently.  There is a similar problem for deletes.


The GoatRope Data Structure
============== 

- Based on a PieceTable, which was one of the things that got early Microsoft Word off the ground.
- A completely separate package from the Braille text editor which is driving it as a specific need.
- Is an actual file, except it has special semantics around seek and write

```go
var _ File = &GoatRope{}
```

The PieceTable can be tested and used on its own, just to track how the edited file is wired together.
The main characteristics are:

- The original file being edited can stay on disk, and is immutable
- A second file that can be on disk or in memory, and is append-only
- A table that points to either the original or mods file, offsets where data is, and how many bytes to use at that offset
- This lets us easily have `Load` to get the initual file, `Insert`, and its inverse `Cut` which both just take a number of bytes
- There is no reason for the table itself to get entangled in file IO, as it is just tracking indices
- The GoatRope wraps around the table with a `File` interface, and the `Original` and `Mods` files only need a file interface
- In this implementation, the Original is on disk, and the Mods is in memory, but they have the exact same interface

```go
	pt := goatrope.PieceTable{}
	pt.Load(100)
	pt.Insert(50)

	checkPieces(t, 2, pt, []goatrope.Piece{
		{true, 0, 100},
		{false, 0, 50},
	})

```

The table rows: `(IsOriginal, Start, Offset)` specify for each piece, which of the two files, where the offset is, and how many bytes.
The algorithm for doing Insert and Cut are rather complicated, but they are completely isolated and tested independent of any file objects.


If you need to be able to insert and delete into large files efficiently, or edit file streams;
you may find this package useful.  You can use the PieceTable on its own, as it is completely separate
from file IO, and tested on its own.  You make GoatRopes backed by files or memory.

A Braille editor is an extreme example of a line-oriented editor.  It makes sense to have the GoatRope
be a file enhanced with navigation, search, and index functions; where the editor that wraps around it
concerns itself completely with ANSI escape sequences and character generation (ie: building characters
out of dots in Braille, should not creep into the GoatRope, but navigation, search, and indexing probably should).
