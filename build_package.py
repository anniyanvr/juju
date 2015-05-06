#!/usr/bin/python
"""Script for building source and binary debain packages."""

from __future__ import print_function

from argparse import ArgumentParser
from collections import namedtuple
import os
import sys


SourceFile = namedtuple('SourceFile', ['sha256', 'size', 'name', 'path'])


def parse_dsc(dsc_path, verbose=False):
    files = []
    with open(dsc_path) as f:
        content = f.read()
    there = os.path.dirname(dsc_path)
    found = False
    for line in content.splitlines():
        if found and line.startswith(' '):
            data = line.split()
            data.append(os.path.join(there, data[2]))
            files.append(SourceFile(*data))
        elif found:
            # All files were found.
            break
        if not found and line.startswith('Checksums-Sha256:'):
            found = True
    return files


def build_binary(dsc_path, location, verbose=False):
    source_files = parse_dsc(dsc_path, verbose=verbose)
    return 0


def main(argv):
    """Execute the commands from the command line."""
    exitcode = 0
    args = get_args(argv)
    if args.command == 'binary':
        exitcode = build_binary(args.dsc,  args.location, verbose=args.verbose)
    return exitcode


def get_args(argv=None):
    """Return the arguments for this program."""
    parser = ArgumentParser("Build debian packages.")
    parser.add_argument(
        "-v", "--verbose", action="store_true", default=False,
        help="Increase the verbosity of the output")
    subparsers = parser.add_subparsers(help='sub-command help', dest="command")
    bin_parser = subparsers.add_parser('binary', help='Build a binary package')
    bin_parser.add_argument("dsc", help="The dsc file to build")
    bin_parser.add_argument("location", help="The location to build in.")
    args = parser.parse_args(argv[1:])
    return args


if __name__ == '__main__':
    sys.exit(main(sys.argv))
