# fen2image

cli tool that generates chess board png images from [FEN](https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation) description :

    rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1

The icons come from [WikiMedia](https://commons.wikimedia.org/wiki/Category:PNG_chess_pieces/Standard_transparent)

## Usage

Will be detailled once the core functionnality is developped

## Todo

board
 -> write tests
 -> convert board fromFEN
 -> validate FEN input expression with regex
options
 -> handle different cell size (-> resize icons)
 -> output filename
 -> cell colors
renderer
 -> PNG Renderer
 -> SVG Renderer ?
 -> resize, invert board (view from blacks), draw row/column names