# plotholes, plotholes, plotholes

*This text is a collection of stubs, many of which require more thought and experimentation to grow into a useful form.*

When dealing with timetravel, plotholes are more a question of 'when' than 'if'. When writing branching narrative, extreme care must be taken to avoid inadvertent resurrections, personality changes, and causality failures of all kinds. When dealing with time travel within a branching narrative - causality problems are a basic inevitability.

The problem is that humans are bad at dealing with complexity. Success in complex work is predicated on the possibility of abstraction, black-boxing, and demarcation - a whole host of techniques that allow humans to concentrate on a specific problem without having to think of all the other problems of a given project.

By necessity, programmers have become extremely capable at this. A large codebase is a little like a novel that is millions of pages long, worked on by hundreds of people, some of whom are retired, that explodes if there are any logical inconsistencies.

Writing a novel like this would be straightforwardly impossible without formal techniques and good tools.

With a Lovecraftian timetravel narrative, we have a much easier task. Weird stuff is expected to happen. Dead characters are expected to appear again, personality switches are the norm.

Nonetheless, it's nice if these things are the product of choice, and not accident. So, the question is, are there insights from the world of software engineering that will allow us to make the complexity of a text somewhat more tractable?

## approaches

### DRY

(Don't Repeat Yourself.)

This is basically the principle that each piece of information should have a singular canonical representation. This is problematic for fiction, since the representation of an object typically continues through the text. 

*Perhaps this could be achieved after the fact? A Regexp could be used to collate the representations of a character in chronological order so they could be checked*

### Composition

Composition typically talks about building more complex objects out of simpler ones. It's a basic idea of programming, at its most simple, the ';' operator in C.

In natural language, this tends to get very repetitive. Dwarf Fortress constructs its descriptions by composition of descriptions of elements. It's also hard to judge what is a self-contained textual element, and what isn't.

*Maybe at a larger scale? Rather than composing characteristics to build a character, you could compose plot beats to construct a story?*

### Modularity

Stories can be insulated from one another. Vignettes and snapshots can be written such that they don't cause side-effects.

## Tools

"Never send a human to do a machines' job"

Some things that are difficult for humans are very simple for machines. Amongst them - checking conformity to simple rules. This means that any desirable trait of a text that can be described as a simple rule can be turned into an automated test.

### memory management.

If a character dies, you could 'free' the memory associated with it, then consider it an error to associate verbs with this character?

If a character does something in one timeline, it should be an error to refer to that action in another timeline.

Using a concept of ownership, it would be possible to write an error checker that gives warnings about likely plot inconsistency.

*This is obviously problematic.*

### race condition checker


