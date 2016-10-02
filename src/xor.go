package main

import (
   nn "neuralnetwork"
   "math/rand"
   "fmt"
)

func __round__ (x float64) float64 {
   if x > 0.5 {
      return 1.0
   }
   return 0.0
}

func main () {
   nn.RandomSeed()
   n := nn.NewNeuralNetwork()
   hidden := 12
   n.AddLayer(nn.NewLayerLinear(1, 2, hidden, 1, 0))
   n.AddLayer(nn.NewLayerActivation(1, hidden, "sigmoid"))
   n.AddLayer(nn.NewLayerLinear(1, hidden, 1, 1, 0))

   /*
      [0 0  0]
      [0 1  1]
      [1 0  1]
      [1 1  0]
    */
   data := nn.NewSimpleMatrix(4, 3)
   data.Data[0][0] = 0.0
   data.Data[0][1] = 0.0
   data.Data[0][2] = 0.0
   data.Data[1][0] = 0.0
   data.Data[1][1] = 1.0
   data.Data[1][2] = 1.0
   data.Data[2][0] = 1.0
   data.Data[2][1] = 0.0
   data.Data[2][2] = 1.0
   data.Data[3][0] = 1.0
   data.Data[3][1] = 1.0
   data.Data[3][2] = 0.0

   error := 0
   for i := 1; i <= 20000; i++ {
      k := rand.Intn(4)
      n.Fit(data.Window(0, k, 2, 1), data.Window(2, k, 1, 1), 0.1)
      if data.Row(k).Data[0][2] != __round__(n.Predict(data.Window(0, k, 2, 1)).Data[0][0]) {
         error ++
      }
      if i % 1000 == 0 {
         fmt.Printf("error:%.2f%%\n", float64(error)/1000.0 * 100.0)
         error = 0
         fmt.Println(data.Row(k).Data[0][0], "xor", data.Row(k).Data[0][1], "=", data.Row(k).Data[0][2], "  [A]", n.Predict(data.Window(0, k, 2, 1)).Map(__round__) )
      }
   }
}
