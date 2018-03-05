#Imaginary Society

 So the idea was to present society as a kind of graph:


						      		      /`Something else
    Entry Somethign --(a)	      		(c)-- Somewhat Something--Something else
    		        \	     	      /			      \_Something else
    		         \	             /				 /``Something else
      Somebody somebody--(b)--  Alson Beric   --(d)-- Julius Simpleton-- Something else
    				  	     \		  		 \__Something else
    				    	      \			
    				      	       (e)--   Alex Grovenor   -	


I think this graph could do double-duty, as a kind of graph editor. The entities in the game world essentially form a graph, with relationships (a - e) that are weighted and gated. One way to manage entity-decision making would be through a graph - entities could have a goal (depending on their personality), then map a path to that goal using the worldgraph - that is, the representation of the totality of entities relations. Basically, some kind of GOAP (Goal-oriented-action planning).

This kind of graph would have to be hand made. For a human, it's obvious you can use a gun to rob a store. For a computer, it isn't. By using a combination of gates (does the entity have a gun?) and weights (is it difficult to rob a store?) an entity could plot a sequence of actions that would lead to a specific goal. All entities would share the same graph, but not the same goals and starting positions.





