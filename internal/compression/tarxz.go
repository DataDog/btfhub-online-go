package compression

import (
	"archive/tar"
	"bytes"
	"io"

	"github.com/rotisserie/eris"
	"github.com/ulikunitz/xz"
)

// General explanation about tar.xz https://linuxize.com/post/how-to-extract-unzip-tar-xz-file/

// ExtractFileTarXZ extracts the file from the tar.xz.
// The input reader expects to get a tar.xz content with a single file in it.
func ExtractFileTarXZ(inputReader io.Reader) ([]byte, error) {
	xzReader, err := xz.NewReader(inputReader)
	if err != nil {
		return nil, eris.Wrap(err, "failed creating xz reader")
	}

	tarReader := tar.NewReader(xzReader)

	for {
		header, err := tarReader.Next()

		switch {
		// if no more files are found return
		case err == io.EOF:
			return nil, nil

		// return any other error
		case err != nil:
			return nil, eris.Wrap(err, "failed iterating over the given tar")

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// if it's a file - create it
		if header.Typeflag == tar.TypeReg {
			buf := &bytes.Buffer{}
			if _, err := io.CopyN(buf, tarReader, header.Size); err != nil {
				return nil, eris.Wrapf(err, "failed writing compressed file content %q", header.Name)
			}
			return buf.Bytes(), nil
		}
	}
}
