[![GO](https://img.shields.io/badge/Go-1.18.2-lightgrey.svg?logo=go)](https://go.dev)
![](https://img.shields.io/badge/License-MIT-blue.svg)
![](https://img.shields.io/badge/build-passing-limegreen.svg)

# A* Search Algorithm

Simple [Go](https://go.dev/) implementation of the [A* search algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm) using a minimal [binary heap](https://en.wikipedia.org/wiki/Binary_heap). The heap's structure ensures that the top element always points to the next neighbor for the shortest path. This project is an extension from an assignment dealing about search algorithms that I initially wrote in [C](https://en.wikipedia.org/wiki/C_(programming_language)).

## 🌱 Project source code

> **Warning**<br>
> Before attempting to run the project, make sure you have all of the files shown in the tree structure below.

```
a-star/
├── libraries/
│   ├── grid/
│   │    └── grid.go
│   └── heap/
│        └── heap.go
├── go.mod
├── main.go
└── Makefile
```

## 🏃🏽 Running the project

Typing `make run` in the terminal will run the project. It should display something like this:

```
valeroy@valentins-macbook-pro a-star % make run
go run main.go

Grid size  : 20x20
Wall units : 50
Start      : {0 0}
End        : {19 19}

s#               #   
..       #           
 ..     #          # 
  .#            #    
  .......       # ## 
       #. ##  ##     
        .  #         
  # #  #.            
 #    # .  #         
  ##    .. # #       
 # #  #  ..   ##     
          ..##       
     #   ##..      # 
 #          ...  #   
#   #        #. #    
        #    #..     
        #      ..    
  #             ..   
                 ..# 
     ##           .e 
                     
Time elapsed: 526.608µs
```