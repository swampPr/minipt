# MinIPT

## A simple, bare-bones image processing toolbox for image conversions, image blurring, setting filters, horizontal and vertical flips and more.

#### Requirements:
- Have [Go](https://go.dev) installed

### Project Setup:

- `git clone git@github.com:swampPr/minipt.git`
- `cd ./minipt`
- `go build`
- `go install`
- `minipt --help`



To view help and examples on a certain subcommand run:

`minipt [SUBCOMMAND] --help`

### Features:
- Turn an image into ASCII art:
  - `minipt ascii ./image.png`
- Turn an image into ANSI art:
  - `minipt ansi ./image.png`
- Resize an image:
  - `minipt resize --bilinear --width 160 --height 160 ./in.png out.png `
- Apply filters to an image:
  - `minipt filter --gray ./in.png ./out.png`
- Convert images to another format:
  - `minipt convert --quality 75 --lossless ./in.png ./out.jpeg`


You can view all subcommands with `minipt --help`

#### Built with Cobra

