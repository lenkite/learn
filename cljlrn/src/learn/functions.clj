(ns learn.functions
  (:require [clojure.io :as io])
  (:import (java io.File net.URL)))


; TODO: Multi-Methods

; *multi-method*

; to-url uses the built-in class function, which will return the underlying
; type of the value as its dispatch function.
(defmulti  to-url class)
(defmethod to-url io.File [f] (.toURL (.toURI f))) (defmethod to-url net.URL [url] url)
(defmethod to-url java.lang.String [s] (to-url (io/file s)))


; *apply*

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
(map vector [:a :b] [:c :d])
;;=> ([:a :c] [:c :d]))

; *partial*
; called partial since it partially fills in the arguments for an existing
; function producing a new function of fewer arguments in the process
; aka currying.

; we know + takes varargs
; currying can ofcourse be done manually. Ex:
(defn my-inc [n] (+ 1 n))
; howweer we can also use partial like the below:
(def my-inc (partial + 1)) 
; note that first formused defn becuase we had a param vector and body, while
; the second form used plain def since we simply bound a value to my-inc

; Further examples
(def dracula {:title "Dracula" :author "Stoker" :price 1.99 :genre :horror})

; The cheaper-than function below takes advantage of Clojure’s truthy logic and
; return nil when the book fails the test, and the book map itself—which is
; truthy—when it passes.
(defn cheaper-than [max-price book]
  (when (<= (:price book) max-price) book))

; variants of cheaper-than using partial
(def cheap? (partial cheaper-than 9.99))
(def crazy-cheap? (partial cheaper-than 1.00))
(defn horror? [book] (when (= (:genre book) :horror) book))

; *complement*
; compoment wraps the function that you supply with a call to not producing a
; new function that is the complement of the original
; Ex

(defn adventure? [book]
  (when (= (:genre book) :adventure) book))

(adventure? dracula); => nil

; can be done manually like:
(defn not-adventure? [book] (not (adventure? book)))
(not-adventure? dracula); => true
; but can be done easily like:
(def not-adventure? (complement adventure?))
(not-adventure? dracula); => true

; *every-pred*
; Higher order function that combines predicate functions into a single function that ands them together.


(def cheap-horror? (every-pred cheap? horror?))
(cheap-horror? dracula) ; => true

; TODO: *lambdas, functional literals* 


; *loop and recur*
; loop works with recur. When it hits a recur inside the body of a loop, Clojure
; will reset the values bound to the symbols to values passed into recur and
; then recursively reevaluate the loop body.

(def books
  [{:title "Jaws"  :copies-sold 2000000}
   {:title "Emma"  :copies-sold 3000000}
   {:title "2001"  :copies-sold 4000000}])

(defn sum-copies [books]
  (loop [books books total 0]
    (if (empty? books)
      total
      (recur (rest books) (+ total (:copies-sold (first books)))))));

(sum-copies books)
(apply + (map :copies-sold books))

; TODO: map 


; Pre and Post Conditions

; To set up a :pre condition just add a map after the parameter vector and
; before the body expression - aa map with a :pre key. The value should be a
; vector of expressions. You will get a runtime exception if any of the
; expressions turn out to be falsy when the function is called
(defn publish-book [book] {:pre [(:title book)]} 
  (print book)
 ;(ship-book book)
 )

; The below ensures that books have both title and author
(defn publish-book [book]
  {:pre [(:title book) (:author book)]} 
  ;(print-book book)
  ;(ship-book book)
  )

; You can also specify a :post condition which lets you check on the value
; returned by the function. Value is available as placeholder %

(defn publish-book [book]
  {
   :pre [(:title book) (:author book)]
   :post [(boolean? %)]
   } 
  ;(print-book book)
  ;(ship-book book)
  ;true)
  0)

(publish-book {:title "bingo" :author "Maddy"})



