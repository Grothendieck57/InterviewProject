# InterviewProject
#GoLang Version 1.18 used

Guys, here is the source code along with any constituent files needed to compile properly. 

First off, I want to thank everyone that has been involved in this process. 
Regardless of the decision that is made based on the work I've produced here, I've learned many valuable lessons throughout this interview process. 
Everything from my conversations to solving the problems has been a true joy, and there is little doubt in my mind about the quality of CommentSold as a company.

On to explaining the files and what is going on.

-interviewPost.go : This file houses the source code for Part 2 of the interview spec supplied to me. It will use data.txt as an input file.
    i) data.txt : This is a text file that contains the data supplied via the project spec.
    
-interviewGet.go : This is my first take on part 1 of the interview spec. However, I don't think the solution is logically complete. "
                    While spidering through the EP via UUIDs seems like the right approach, this code skips quite a few of the additonal
                    UUIDs found in the Array items of each customer. I find it hard to believe that some of those UUIDs don't correspond
                    to other unique, unvisited customers on the endpoint.

-interviewGetV2.go : I am convinced this approach is more complete, but it is TERRIBLY inefficient. Clearly, there is a better train of logic--perhaps 
                      a different logical appraoch to traversing the EP--that would get me a solution in a timely manner. Eitehr way, the idea here is 
                      that I don't leave any of the "spare" array UUIDs unvisited, unless they are found to be duplicate customers. Unforunately, the
                      arrays used within the nested for loops are large--causing incredible cost to runtime. There's a good chance I might be overthinking
                      things here. Sorry.
                      
                      
