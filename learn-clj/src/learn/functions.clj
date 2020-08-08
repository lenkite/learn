(ns learn.functions)

; TODO: M
; TODO: Multi-Methods
; Loop and recu


; The apply function

; apply basically unwraps a sequence and applies the function to them as
; individual arguments.  You use apply to convert a function that works on
; several arguments to one that works on a single sequence of arguments. You
; can also insert arguments before the sequence
; The clojure doc for apply sadly never talks about this vital unwrapping
; apply can be visualized thinking about "unrolling" or "spreading" arguments from a list to call a function.

(str (reverse "derp"))
;; => "(\"p\" \"r\" \"e\" \"d\")"

(apply str (reverse "derp"))
;; => "pred"

; when you call reverse on a string it returns a sequence of character string in the reverse order.
; Str is interesting because it can be called on one thing or more things. If
; called on just one thing then it stringifies the thing and if passed many
; arguments it will stringify them and then concatenate them. So, this is why
; the first incorrect version returns what it does. It sees the list of
; character strings as one thing- a list, and it then nonchalantly returns the
; stringified version of the whole list. The second example, on the other hand,
; is saying take the str function and apply it to all the arguments.


; classic example to transpose a matrix from clojuredoc
; http://clojuredocs.org/clojure.core/apply#example-542692cdc026201cdc326d4d
(apply map vector [[:a :b] [:c :d]])
;;=> ([:a :c] [:b :d])

; The one inserted argument here is vector. So the apply expands to
; Basically the previous apply expands to:
(map vector [:a :b] [:c :d])




