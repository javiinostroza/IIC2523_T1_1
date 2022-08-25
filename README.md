# IIC2523_T1_1

- Benjamín Yon Bazán
- Javiera Inostroza Ríos


## Comentarios sobre elección de algoritmo de busqueda
- Para el problema que estamos resolviendo, independiente de si es DFS o BFS la complejidad temporal es O(N * M), ya que tendremos que visitar todos los nodos al menos una vez, con el algoritmo que sea. Escogimos DFS porque nos fue más facil de implementar que BFS y usamos la versión recursiva ya que es más sencilla que la iterativa.

- Para el caso de complejidad espacial (memoria auxiliar utilizada), fue resuelto usando una matriz de visitados que ocupa memoria auxiliar O(N * M) y a esto le sumamos las llamadas recursivas que en CANTIDAD son O(N * M) y esto usa espacio en el stack.


- Existen versiones del algoritmo que reutilizan la matriz cargada, para añadir un estado extra que indica que un nodo fue visitado. Con esto no se necesita la matriz de visitados y se reduce la memoria auxiliar a utilizar.
- En el caso iterativo podríamos reducir la memoria utilizada por las recursiones a O(N * M) elementos en la cola/lista de elementos por visitar (distinto a llamadas recursivas que usan stack)
- Para nuestro problema la memoria auxiliar no varía para BFS ni DFS.
