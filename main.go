/*
* @project     gcolor
* @package     gcolor
* @author      @jeffotoni
* @size        21/08/2017
*
 */

package gcolorfake

// Sufix - Is added to the end of each color
const Sufix = "\033[0m"

// Black - Return text with black color
func BlackCor(msg string) string {
	return "\033[0;30m" + msg + Sufix
}

// Red - Return text with Red color
func RedCor(msg string) string {
	return "\033[0;31m" + msg + Sufix
}

// Green - Return text with black color
func GreenCor(msg string) string {
	return "\033[0;32m" + msg + Sufix
}

// Yellow - Return text with Yellow color
func YellowCor(msg string) string {
	return "\033[0;33m" + msg + Sufix
}
