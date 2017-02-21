package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var defaultBoard = [8][8]byte{
	{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
	{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}}

func contains(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isValidFen(fen string) bool {
	var linePattern = "[1-9rnbqkpRNBQKP]+"
	var fenRegex = fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s w|b (K?Q?k?q?)|- ([a-h][1-8])|- \\d+ \\d+", linePattern, linePattern, linePattern, linePattern, linePattern, linePattern, linePattern, linePattern)
	var isValid, _ = regexp.MatchString(fenRegex, fen)
	return isValid
}

func getBoardLine(fenElement string) ([8]byte, error) {
	var line = [8]byte{}
	var pos = 0
	for _, cellValue := range fenElement {
		if pos >= 8 {
			return line, errors.New(fmt.Sprintf("Oops: %s is not a valid fen element", fenElement))
			break
		}
		if contains(validPieces, byte(cellValue)) {
			line[pos] = byte(cellValue)
			pos = pos + 1
			continue
		} else {
			pos = pos + int((int(cellValue) - int('0')))
		}
	}

	return line, nil
}

func getBoardFromFen(fen string) ([8][8]byte, error) {
	if false == isValidFen(fen) {
		return defaultBoard, errors.New("Invalid FEN !")
	}
	var splits = strings.Split(fen, " ")
	var lines = strings.Split(splits[0], "/")

	var board = [8][8]byte{}
	for y := 0; y < 8; y++ {
		var err error
		board[y], err = getBoardLine(lines[y])
		if err != nil {
			return defaultBoard, err
		}
	}

	return board, nil
}

func isChessPieceOrPawn(piece byte) bool{
	return piece != ' ' && contains(validPieces, piece)
}
