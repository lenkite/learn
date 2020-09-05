(ns learn.collections)


; Various map functions

; *update* function
; updates a value in map, takes m(map, k(key) and f(function) as arguments followed by optional args)
; f is a function that will take in old value and args and return the new value.
; update will return a new map
(def book {:title "Emma" :copies 1000})

; (update [m k f & args])
(update book :copies inc) ;=   {:title "Emma", :copies 1001}


; *update-in* function
(def by-author
  {:name "Jane Austen"
   :book {:title "Emma" :copies 1000}})

; (update-in [m ks f & args]) ; update-in takes in key-sequence vector instead
; of plain key. creates levels if they don't exit
(update-in  by-author [:book :copies] inc)

; *assoc-in* function
; associates a new value in a map for a given key sequence 
; (assoc-in [m ks v])

; Example  using Ring and Jetty. 
; In Ring requests and responses are maps. A handler is a function that takes a
; request map and returns a response map
(defn handler [ request]
  {:status 200
   :headers {"Content-Type" "text/html"}
   :body "Hello from your web-app"})

; you can start web-app by passing the handler function to rings run-jetty
;(defn -main [] (jetty/run-jetty handler {:port 8080}))

; ring apps use middleware which are functions that take a handler function (and other arguments) as
; parameters and return a new handler function. Can be used to layer
; additional features onto handlers
(defn log [msg value] 
  "Logs message and value. Returns the value"
  (println msg value) value)

(defn wrap-log 
  "Returns a middlware function that logs the response"
  [msg handler] 
  (fn [request] 
    (log msg (handler request))))

(defn wrap-content-type
  "Middlware function that Returns a function that sets the response content type"
  [handler content-type]
  (fn [request] 
    (assoc-in (handler request) [:headers "Content-Type"] content-type))) ;assoc-in will create levels if they don't exist



