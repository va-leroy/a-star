# A* Search Algorithm Go Implementation

[![GO](https://img.shields.io/badge/Made%20with-Go%201.18.2-lightgrey.svg?logo=go)](https://go.dev)

## 🎯 Introduction

A very simple implementation of the [A* search algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm) in [Go](https://en.wikipedia.org/wiki/Go_(programming_language)). This implementation is a little bit special, indeed it uses a [priority queue](https://en.wikipedia.org/wiki/Priority_queue) called P and a minimum [binary heap](https://en.wikipedia.org/wiki/Binary_heap) called Q. 

Thanks to the heap, we always know which neighbor to go to in order to get the shortest path, since this neighbor is always on top of Q. I also programmed this algorithm in [C](https://en.wikipedia.org/wiki/C_(programming_language)) this year, but the use of pointers and the very practical [garbage collector](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) of Go, make me prefer this language from now on, which I find just as fast for what I want to do.

## 🧬 Tree structure

> **Warning**<br>
> Before trying to launch the project, please make sure you have **all the files** in the tree structure shown below.

```sh
a-star/
├── libraries/
│   ├── grid/
│   │    └── grid.go
│   └── heap/
│        └── heap.go
├── .gitignore
├── Makefile
├── LICENSE.md # Optional
├── README.md # Optional
├── go.mod
└── main.go

3 directories, 8 files
```

## ⚡️ Start the project

Here is the command to start the project:

```sh
make run # Starts the project
```

> **Note**<br>
> This command will not display anything particular in the console except the parameters used for the creation of the graph and the walls, but it will create a `grid.txt` file displaying the path of the algorithm and thus the shortest path.

```
valeroy@valentins-macbook-pro a-star % make run
go run main.go
[*] Grid size: 20
[*] Number of walls: 50
[!] Path found
[!] Time elapsed: 747.374µs

valeroy@valentins-macbook-pro a-star % cat grid.txt
   s                 
   .         #       
  #.     #           
   .    # #  #     # 
   . #  #  #         
   .  ##             
#  .        #        
   ...     #    #    
    #.   #     #     
 #   ...   #    #    
      #.       #     
    #  ..#  #  #   # 
        .#   #       
       #....      ## 
  #   # # #.    #    
           ..   # #  
#   #       ..       
           # ..      
         # #  .#     
         #   #...    
                .e    
```

Here is the command to clean the project:

```sh
make clean # Removes the `grid.txt` file
```

## 🦾 Benchmarking

I will insert statistics and benchmarking tests in this section later on.

## 📖 License

This project is free and open-source, licensed under the [MIT License](LICENSE.md).