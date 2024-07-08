### 1. Social Networks

#### Friendship Relationships:
In platforms like Facebook, each user can be considered a node, and their friendships as edges. An adjacency list efficiently stores and processes these relationships.

#### Followers/Following:
In Twitter, an adjacency list can represent who follows whom, making it easier to recommend followers or find mutual followers.

### 2. Running the Program

To run the program, follow these steps:

1. Enter an edge (from to) or type 'done' to finish: Alice Bob
   - Adding edge from Alice to Bob
   
2. Enter an edge (from to) or type 'done' to finish: Alice Charlie
   - Adding edge from Alice to Charlie
   
3. Enter an edge (from to) or type 'done' to finish: Bob David
   - Adding edge from Bob to David
   
4. Enter an edge (from to) or type 'done' to finish: Charlie David
   - Adding edge from Charlie to David
   
5. Enter an edge (from to) or type 'done' to finish: done

Final graph:
- Alice: [Bob Charlie]
- Bob: [Alice David]
- Charlie: [Alice David]
- David: [Bob Charlie]
