# roll
Roll is a simple dice rolling simulator written in Go. It takes a single integer as a command line agrument and then prints that many six-sided dice to stdout. It is a slight expansion of a demonstration I made for a stackoverflow question. I had fun writing it to answer the question so I decided to expand it out a little and push it.

./roll 5 looks something like this in the terminal:

    ┌───────┐  ┌───────┐  ┌───────┐  ┌───────┐  ┌───────┐  
    │ ∙     │  │ ∙     │  │       │  │ ∙     │  │ ∙   ∙ │  
    │   ∙   │  │   ∙   │  │ ∙   ∙ │  │   ∙   │  │   ∙   │  
    │     ∙ │  │     ∙ │  │       │  │     ∙ │  │ ∙   ∙ │  
    └───────┘  └───────┘  └───────┘  └───────┘  └───────┘ 

The original quetion is at:

https://stackoverflow.com/questions/36220839/print-multi-line-strings-on-same-line

