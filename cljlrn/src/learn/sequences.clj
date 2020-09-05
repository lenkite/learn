; A seqable is something that the seq function can turn into a sequence.


;  The cons function stands for “construct” and takes two arguments, an item
;  and a sequence. It returns a new sequence created using the item as its
;  first and the sequence as its rest. A sequence created by cons is known as a
;  “cons cell”—a simple first/rest pair. Sequences of any length can be
;  constructed by chaining together multiple cons cells.

(cons 4 '(1 2 3)); => (4 1 2 3))

; The conj function is similar to cons and stands for “conjoin.” The main
; difference from cons is that (if possible) it reuses the underlying
; implementation of the sequence instead of always creating a cons cell. This
; usually results in sequences that are more efficient. Whether the new item is
; appended to the beginning or end depends on the underlying representation of
; the sequence. Unlike cons, conj takes a sequence as its first parameter, and
; the item to append as the second

; For both conj and cons, if you supply nil in place of the sequence, it constructs a sequence containing only one item, the one you specified.
(cons 1 nil) ; => (1)
(conj nil 1) ; => (1)

;This is used in another common Lisp idiom, constructing a list recursively
;using cons or conj. The following function demonstrates recursively
;constructing a sequence of all the integers from 1 to the provided parameter

(defn make-int-sequence [max]
  (loop [intseq nil n max]
    (if (zero? n)
      intseq
      (recur (cons n intseq) (dec n)))))

(make-int-sequence 10)

; The first/rest architecture of sequences is the basis for another extremely important aspect of Clojure sequences: laziness
; Laziness is made possible by the observation that logically, the rest of a
; sequence doesn’t need to actually exist, provided it can be created when
; necessary. Rather than containing an actual, concrete series of values, the
; rest of a lazy sequence can be implemented as a function which returns a
; sequence.
; In a lazy sequence, calling rest actually calculates and instantiates the new
; sequence, with a freshly calculated value for its first and updated
; instructions on how to generate still more values as its rest.


; To see a lazy sequence at work, consider the map function. The map function
; is an extremely important sequence manipulation tool in Clojure. It works by
; taking a sequence and a function as arguments, and returns a new sequence
; which is the result of applying the supplied function to each of the values
; in the original sequence. 

(map #(* % %) '(1 2 3 4 5 6 7))

; To see the internal workings of the lazy sequence, let’s add a side effect to
; your square function, so you can see when it’s being executed (normally, side
; effects in functions provided to map are not a great design practice, but
; here they will provide insight into how lazy sequences work). 

(defn square [x] 
  (println (str "Processing: " x)) (* x x))

; You now have a symbol map-result which is, supposedly, bound to a sequence of
; the squares. However, you didn’t see the trace statement. square was never
; actually called! map-result is a lazy sequence. Logically, it does contain
; the squares you expected, but they haven’t been realized yet. It’s not a
; sequence of squares, but a promise of a sequence of squares. You can pass it
; all around the program, or store it, and the actual work of calculating the
; squares is deferred until it is required.

(map square '(1 2 3 4 5))
; The reason why the code is so ugly is that the println calls are being called in the middle of printing out the results

(def map-result (map square '(1 2 3 4 5)))
(nth map-result 2)
; Processing:1
; Processing:2
; Processing:3
; 9
(nth map-result 2); => 9 ; does not process again, caches

(println map-result)

; partition, which chops up a big sequence into a sequence of smaller sequences, enabling you to take a flat vector
(def titles-and-authors ["Jaws" "Benchley" "2001" "Clarke"])
(partition 2 titles-and-authors)
