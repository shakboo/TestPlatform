#coding=utf-8

import csv
import sys

class changeFileFormat():
    def __init__(self):
        pass

    def main(self, path, rawFormat, targetFormat, targetPath, splitStr=None, sourceSplitStr=None):
        with open(path, 'r') as f: 
            content = f.read()
            newContent = content.replace(self._get_seq(rawFormat, splitStr=sourceSplitStr), self._get_seq(targetFormat, splitStr=splitStr))
        with open(targetPath, 'w') as f:
            f.write(newContent)

    def _get_seq(self, format, splitStr=None):
        if format == "csv":
            seq = ','
        elif format == "tsv":
            seq = '\t' 
        elif format == "xsv":
            seq = splitStr
        elif format == "txt":
            seq = " "
        return seq

if __name__ == '__main__':
    c = changeFileFormat()
    data = c.main(*sys.argv[1:])
