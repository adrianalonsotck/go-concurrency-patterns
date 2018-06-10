# Barrier Concurrency Pattern

En concurrencia, una barrera es un tipo de método de sincronización. Una barrera para un grupo de subprocesos significa que cualquier sub-proceso debe detenerse en este punto y no puede continuar hasta que todos los subprocesos lleguen a esta barrera. La barrera es un patrón muy común, especialmente cuando tenemos que esperar más de una respuesta de diferentes fuentes antes de permitir que el programa continúe. Descubramos en que consiste el Barrier Concurrency Pattern.

## Descripción

Imagina una situación en la que tenemos una arquitectura de software orientada a microservicios. En esta un servicio necesita componer su respuesta (como si se tratara de API Gateway) fusionando las respuestas de otros tres microservicios. Aquí es donde el patrón puede realmente ayudarnos.

Nuestro patrón podría ser un servicio intermedio que bloqueará su respuesta hasta que se haya compuesto con los resultados devueltos por uno o más goroutines diferentes. Tradicionalmente, podríamos implementar este patrón con un Lock, pero en Go podemos usar un canal sin buffer para llevarlo a cabo.

## Objetivo
		 	 	 		
Fiel a su nombre, el patrón intenta detener una ejecución por lo que no termina antes de que estén resultos sus subprocesos para finalizar. Los objetivos del patrón de Barrier son los siguientes:

- Componer el valor de un tipo con los datos que provienen de N goroutines.
- Controlar la completitud de cualquiera de esos datos entrantes para que no se devuelvan datos incoherentes. No se desea un resultado parcialmente llenado porque una de las tuberías ha devuelto un error.

## Ejemplo: HTTP Aggregator

Para nuestro ejemplo, vamos a escribir una situación muy típica en una aplicación de microservicios, una aplicación que realiza N llamadas HTTP GET y las une en una única respuesta que se imprimirá en la consola.

Nuestra pequeña aplicación debe realizar cada solicitud en un goroutine diferente e imprimir el resultado en la consola si ambas respuestas son correctas. Si alguno de ellos devuelve un error, imprimimos el error.

## Resúmen
El patrón Barrier abre la puerta de la programación de microservicios con su naturaleza composición. El patrón no sólo es útil para realizar solicitudes de red; también podríamos usarlo para dividir algunos problemas complejos en simples tareas que ejecutará cada goroutine.

Por ejemplo, una operación costosa podría dividirse en algunas operaciones más pequeñas distribuidas en diferentes Goroutines para maximizar el paralelismo y lograr un mejor rendimiento.
