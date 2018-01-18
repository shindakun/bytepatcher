# bytepatcher

<p align="center">
  <img style="float: right;" src="assets/bytepatcher.png" alt="bytepatcher gopher"/>
</p>

Sometimes you'll run into a situation where you need to patch a binary file. bytepatcher is a small utility writen in Go which can be used when you know the signature and patch hex.

bytepatcher was created as a learning project.

## Installation

If you are a Go user you can run `go install github.com/shindakun/bytepatcher`. Prebuilt binaries are not currently provided.

## Usage

$ ./bytepatcher
  -in string
        input filename
  -out string
        output filename
  -patch string
        patch hex as string
  -sig string
        signature as string

## Usage example

For a very basic example, say we need to change a JE (74) to a JMP (EB) in an executable file.

```bash
$ bytepatcher -in file.exe -out patched.exe -sig 83F8057441 -patch 83F805EB41
2018/01/17 13:20:00 Signature: 83f8057441
2018/01/17 13:20:00 Patch: 83f805eb41
2018/01/17 13:20:00 Found signature at 0x5e3!
2018/01/17 13:20:00 Wrote 5 bytes!
2018/01/17 13:20:00 Saved output!
```

You'll note we used a long signature and patch, this is because we wanted to patch one specific location in the file. If we had used a shorter signature we would potentially patch in the wrong spot. Also it is worth noting that currently we're only able to patch a single location. So, even though we're writing 5 bytes we're really only patching 1.
