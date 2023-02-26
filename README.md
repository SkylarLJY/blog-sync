# Blog sync
A cmd line tool to sync and reformat markdown files between folders. It creates a new folder with the name of the source file and copy the content into `index.md` int this folder.

It's called blog sync because it reformats the source file to add metadata for blog post display. The reformatting can be disabled. 

The resulting file shoud look like
```
---
title: xxx
date: xxx
description: xxx
---
...content of the source file...
```
## Usage 
Specify `BLOG_SYNC_SRC_DIR` and `BLOG_SYNC_DEST_DIR` as environment variables or in a .env file. 

`-file`: string - required, name of the source markdown file **without** the `.md` extension 
`-destName`: string -  optional, name of the destination folder, if not specified will be the source file name 
`-format`: boolean - optional, true by default (which adds the metadata header)

Example:
`blog-sync -file input -destName output` 
should copy the file `BLOG_SYNC_SRC_DIR/input.md` into `BLOG_SYNC_DEST_DIR/output/index.md` and add a header
```
---
title: "input"
date: "2023-02-26T16:18:58-05:00"
description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure ..."
---
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco 

laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu 

fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
```