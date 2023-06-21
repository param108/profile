# threads

This document attempts to define a thread in the context of `tribist`.

# overview

On twitter, a thread is basically a large document divided into tweets.
They are meant to be read together, although they may appear on the timeline
at different times.

You can add to threads and remove from them at any time.

Usually, the first tweet sets the scene for the rest of the thread.

# A thread as additional knowledge

We propose that a thread is really a kind of footnote.
Some extra information for the curious.

The first tweet, then, is the summary and the title page (as it were).

Reading a thread, does not lose your place in the main timeline.

In this context, parts of the thread appearing at different times on
your timeline seems ridiculous. Instead it is enough to show the 
whole thread reappear on your timeline with a signage saying
"new tweets added" - maybe a red dot.

# Behaviour

In light of this a thread should behave in the following way.

1. In the timeline the first tweet member of a thread (may not be the title tweet),
will be replaced by the title tweet with some marking saying that a new tweet has appeared.
    - How do we define newness ?
2. All other members of the thread are hidden from view.
3. Clicking on the title tweet opens up a parallel timeline to the right of the main timeline
showing the tweets in the thread in sequential order. The tweet right next to the title tweet
on the main timeline should be the original tweet on the main timeline.
4. Threads can be nested. 
5. On a large screen, clicking on a nested thread opens a new page where the thread now is shown 
as the main thread (on the left) and the new thread is on the right. We must keep track of the 
scroll value for all previous threads.
6. threads need not share tags. The connection and sequence will be handled
using a display tag (first line tag `#thread:<thread id>:<sequence number>`).
    - How will we avoid crossing the tweet length ?
         - Only two threads per tweet ?
           You can only have at most two #thread tags per tweet and one of them
           must be sequence number 0 (title thread)
7. On mobile, we just move to new pages with the new thread. We must maintain the scroll
somewhere so that we can seamlessly return.
