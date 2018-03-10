
##UNICODE REPRESENTATION OF GOL

Using unicode, it's possible to represent the game of life on a scale of 8 cells per letter, using the full 256-character range of Unicode braile. This allows for the compact representation of complex patterns without recourse to obnoxious scrolling.

    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
    █⠓⠑⠇⠇⠕⠓⠑⠇⠇⠕⠓⠑⠇⠇	⠓⠑⠇⠇⠕⠓⠑⠇ █
    █	⠓⠑⠇⠇⠕⠓⠑⠇⠇⠕	      	 █
    █⠓⠑⠇⠇⠕⠓⠑⠇⠇⠕⠓⠑⠇⠓⠑⠇⠇⠕⠓⠑⠇⠕	 █
    █			      	 █
    █			      	 █
    █			      	 █
    █			      	 █
    █			      	 █
    █			      	 █
    ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀

(note: the above is composed with real 6- character braille characters, but any combination of dots is allowed in unicode).

For character editing, presumably, we could have two options:
Allow the player to edit with a zoomed-in mode, where they can place "▀" or "▄" characters using a "█" cursor (or similar).

======================
##ALIEN ENTITIES IN PLAINTEXT

A major problem in the terminal-interface conceit of the game is the problem of representing the sinister and alien within really tight graphical constraints.  I think GOL goes some way towards this, but there are additional ways in which we could introduce threat.

1. One idea was to have the GOL sections representing some kind of 'backend' for 
spells - generators which produce resources that the player can then use for spell research, upkeep or casting (something like Master of Magic). In this context, rather than having hard limits on resource usage - it would be thematic to have increasing instability as the player gets close to the breadline.

We could represent this by simply salting the GOL patterns the player has build with random new 'alive' cells. 

Most GoL patterns are robust enough to take small quantities of this without degenerating - but it's also possible for a single alive cell to cause the entire pattern to explode.

Another way to represent threat would to have the 'walls' of the container have a regenerating 'durability' counter, decremented every time a cell hits the wall,
which would be overcome by a high rate of impacts, making the cells breach their containment:

    			⠑⠇⠇⠕⠓
    			⠑⠇⠇⠕⠓⠑⠇⠇⠕⠓
    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄⠇⠕▄▄▄▄▄
    █⠓⠑⠇⠇⠕⠓⠑⠇⠇⠕⠓⠑⠇⠇	⠓⠑⠇⠇⠕⠓⠑⠇ █
    █	⠓⠑⠇⠇⠕⠓⠑⠇⠇⠕	      	 █
    █⠓⠑⠇⠇⠕⠓⠑⠇⠇⠕⠓⠑⠇⠓⠑⠇⠇⠕⠓⠑⠇⠕	 █
    █			      	 █
    █			      	 █
    █			      	 █
    █			      	 █
    █			      	 █
    █			      	 █
    ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀

