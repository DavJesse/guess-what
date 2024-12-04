# guess-it-2

## Overview

**Gues-It-2** is a command-line program designed to predict a range in which the next number in a sequence will fall. The input sequence represents a graph where the line numbers are the x-axis and the numbers themselves are the y-axis. This program employs statistical and mathematical methods to calculate these ranges, balancing accuracy and performance.

The smaller the predicted range, the higher the score, provided the actual number falls within the range.

---

## Features

- **Dynamic Range Calculation**: Determines the range for the next number based on previous inputs.
- **Efficient Performance**: Optimized for handling extensive input data sets.
- **Testing Compatibility**: Designed to integrate seamlessly with external test datasets.

---

## Usage

### Input Format
The program expects a sequence of numbers as standard input, one number per line. For example:

```
189
113
121
114
145
110
```

### Output Format
The output consists of:
1. The input number.
2. A predicted range (lower and upper bounds) for the next number.

Example:

```
$ ./student/script.sh
189
106 272
113
68 264
121
58 224
```
## Code Explanation
The program reads the data from the standard input and predicts the range for the user's next number. The program uses inter-quartile range to eliminate outlier and make accurate guesses. The range is printed to the terminal.

## Technical Details
- The program is written in Go.
- It uses built-in functions to read files and perform mathematical calculations.
- Error handling is included to handle cases where the file cannot be read or the data is invalid.
## Contributing
If you encounter any issues or have suggestions for improvements, feel free to contribute to the project. You can submit bug reports or pull requests on the GitHub repository.
## License
This project is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License). Feel free to use, modify, and distribute the code for personal or commercial purposes.
## Author
[David Jesse Odhiambo](https://david-jesse.vercel.app/)

Apprentice Software Developer, Zone01 Kisumu