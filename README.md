# GO Walk
Compute the path between two words with each step being a single letter mutation of the previous word. 
Valid mutations are adding a single letter anywhere in the word, dropping a single letter, or changing a letter.  
For example the following list would be valid (but not exhaustive) for the word 'toe': 
 * tome
 * to
 * tie

## Requirements
The Gorilla Mux library is required: 
   
    go get -u github.com/gorilla/mux 
    
    
## API

#### GET /words/{word}
find a word in the current dictionary

response: **200 OK**

    {
        "word": string,
        "exists": boolean
    }

#### PUT /words/{word}
add a word to the current dictionary.  
This is an idempotent operation, it will return successfully even if the word already exists in the dictionary (put 
will not add a second instance of the word)

response: **200 OK**

    {}
 
#### GET /routes?start=word1&end=word2
Compute the dictionary walk between the words _word1_ and _word2_ from the url query parameters.

response: **200 OK**

    {
        "start":"word1",
        "end":"word2",
        "path": [
            "step1",
            "step2",
            "step3"
        }
    }
 
 