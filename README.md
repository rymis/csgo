Constructive Solid Geometry implementation in Go
================================================

TODO: add some description

Usage
-----

Function contains several operations. Small example shows what you can do with the models:

```go
func saveStl(fnm string, model *csgo.Model) {
	f, err := os.Create(fnm)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	csgo.SaveStl(f, model)
}

func main() {
	cube := csgo.NewCube(10.0, 10.0, 10.0)
	sphere := csgo.NewSphere(6)

	saveStl("subtract.stl", cube.Subtract(sphere))
	saveStl("union.stl", cube.Join(sphere))
	saveStl("intersection.stl", cube.Intersect(sphere))
	saveStl("hull.stl", cube.Hull(sphere))
}
```

Generated figures are here:

### Subtract

```stl
solid ScadSharp
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -1.066713 5.000000 -5.000000
       vertex -5.000000 0.009602 -5.000000
       vertex -5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.641006 4.641006 -5.000000
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 4.401134 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.103734 5.000000 -5.000000
       vertex 5.000000 5.000000 -5.000000
       vertex 4.641006 4.641006 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -1.066714 -5.000000 -5.000000
       vertex -5.000000 -5.000000 -5.000000
       vertex -2.800385 -2.800385 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -2.800386 -2.800386 -5.000000
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 -0.009602 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 1.066715 -5.000000 -5.000000
       vertex 5.000000 -0.009602 -5.000000
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.427082 4.427082 -5.000000
       vertex 4.641006 4.641006 -5.000000
       vertex 4.796556 4.537072 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.340892 2.114612 -5.000000
       vertex -1.066713 5.000000 -5.000000
       vertex 4.103734 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.340892 2.114612 -5.000000
       vertex 4.103734 5.000000 -5.000000
       vertex 4.641006 4.641006 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.340892 2.114612 -5.000000
       vertex 4.641006 4.641006 -5.000000
       vertex 4.427080 4.427080 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.303836 -3.430388 -5.000000
       vertex 1.066715 -5.000000 -5.000000
       vertex -1.066714 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.303836 -3.430388 -5.000000
       vertex -1.066714 -5.000000 -5.000000
       vertex -1.866927 -3.984721 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.800385 2.800385 -5.000000
       vertex 4.427082 4.427082 -5.000000
       vertex 4.796556 4.537072 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.800385 2.800385 -5.000000
       vertex 4.796556 4.537072 -5.000000
       vertex 5.000000 4.401134 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.800385 2.800385 -5.000000
       vertex 5.000000 4.401134 -5.000000
       vertex 5.000000 0.009600 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.071069 3.725713 -5.000000
       vertex 4.427080 4.427080 -5.000000
       vertex 2.800385 2.800385 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -4.933971 0.093377 -5.000000
       vertex -3.340892 2.114612 -5.000000
       vertex -2.731242 2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -4.933971 0.093377 -5.000000
       vertex -3.981888 1.045458 -5.000000
       vertex -3.197849 -2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 -1.000000
    outer loop
       vertex -4.933971 0.093377 -5.000000
       vertex -3.197849 -2.296101 -5.000000
       vertex -5.000000 -0.009602 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -4.933971 0.093377 -5.000000
       vertex -5.000000 -0.009602 -5.000000
       vertex -5.000000 0.009602 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.395922 2.395922 -5.000000
       vertex 2.800385 2.800385 -5.000000
       vertex 3.340891 2.114612 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 0.000000 3.109171 -5.000000
       vertex 2.071069 3.725713 -5.000000
       vertex 2.800385 2.800385 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 0.000000 3.109171 -5.000000
       vertex 2.800385 2.800385 -5.000000
       vertex 2.395922 2.395922 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -0.521897 -3.805954 -5.000000
       vertex -1.866927 -3.984721 -5.000000
       vertex -2.800385 -2.800385 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -0.521897 -3.805954 -5.000000
       vertex -2.800385 -2.800385 -5.000000
       vertex -2.299216 -2.299216 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.420069 -1.349004 -5.000000
       vertex -2.299216 -2.299216 -5.000000
       vertex -2.800386 -2.800386 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.420069 -1.349004 -5.000000
       vertex -2.800386 -2.800386 -5.000000
       vertex -3.197849 -2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.485281 -1.071069 -5.000000
       vertex -3.981888 1.045458 -5.000000
       vertex -2.731242 2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.485281 -1.071069 -5.000000
       vertex -2.731242 2.296101 -5.000000
       vertex -2.252885 2.438504 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 2.303836 -3.430388 -5.000000
       vertex 0.715963 -3.641432 -5.000000
       vertex 5.000000 -0.009602 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302891 -2.296099 -5.000000
       vertex 3.694280 1.666246 -5.000000
       vertex 5.000000 0.009600 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302891 -2.296099 -5.000000
       vertex 5.000000 0.009600 -5.000000
       vertex 5.000000 -0.009602 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.607687 1.419651 -5.000000
       vertex 2.731235 2.296102 -5.000000
       vertex 3.340891 2.114612 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.607687 1.419651 -5.000000
       vertex 3.340891 2.114612 -5.000000
       vertex 3.694280 1.666246 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.340894 1.686443 -5.000000
       vertex 3.607687 1.419651 -5.000000
       vertex 3.485282 1.071069 -5.000000
    endloop
  endfacet
  facet normal -0.000000 -0.000000 -1.000000
    outer loop
       vertex 3.340894 1.686443 -5.000000
       vertex 3.420069 1.349002 -5.000000
       vertex 2.299215 2.299215 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.340894 1.686443 -5.000000
       vertex 2.299215 2.299215 -5.000000
       vertex 2.395922 2.395922 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.340894 1.686443 -5.000000
       vertex 2.395922 2.395922 -5.000000
       vertex 2.731235 2.296102 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.071067 2.492629 -5.000000
       vertex 2.395922 2.395922 -5.000000
       vertex 2.299215 2.299215 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302888 2.296101 -5.000000
       vertex 3.420069 1.349002 -5.000000
       vertex 3.485282 1.071069 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302888 2.296101 -5.000000
       vertex 3.485282 1.071069 -5.000000
       vertex 3.109172 -0.000003 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.071070 -2.492628 -5.000000
       vertex 0.715963 -3.641432 -5.000000
       vertex -0.521897 -3.805954 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.071070 -2.492628 -5.000000
       vertex -0.521897 -3.805954 -5.000000
       vertex -0.994564 -3.405246 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex -3.420069 -1.349004 -5.000000
       vertex -3.485281 -1.071069 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex -3.485281 -1.071069 -5.000000
       vertex -3.109172 -0.000001 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302891 2.296099 -5.000000
       vertex -2.252885 2.438504 -5.000000
       vertex -2.071066 2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -0.000002 -3.109172 -5.000000
       vertex -0.994564 -3.405246 -5.000000
       vertex -2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.009602 -5.000000
       vertex -5.000000 5.000000 0.886578
       vertex -5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 -0.009601 5.000000
       vertex -5.000000 -5.000000 1.066713
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 4.246193 -0.002599
       vertex -5.000000 3.079603 3.079603
       vertex -5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 4.246193 -0.002599
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 5.000000 0.886578
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.352747 5.000000
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 3.079603 3.079603
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -0.009602 -5.000000
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 -2.299215 -2.299215
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 -2.299216 -2.299216
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 -5.000000 0.886578
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.348922 2.368046
       vertex -5.000000 2.800386 2.800386
       vertex -5.000000 3.079603 3.079603
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.009603 5.000000
       vertex -5.000000 2.352747 5.000000
       vertex -5.000000 3.079603 3.079603
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.009603 5.000000
       vertex -5.000000 3.079603 3.079603
       vertex -5.000000 2.800386 2.800386
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 3.868776 0.994560
       vertex -5.000000 4.246193 -0.002599
       vertex -5.000000 3.641432 -0.715963
    endloop
  endfacet
  facet normal -1.000000 -0.000000 -0.000000
    outer loop
       vertex -5.000000 3.348922 2.368046
       vertex -5.000000 3.768895 1.258452
       vertex -5.000000 2.513672 2.513672
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.348922 2.368046
       vertex -5.000000 2.513672 2.513672
       vertex -5.000000 2.800386 2.800386
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.093323 4.934014
       vertex -5.000000 2.800386 2.800386
       vertex -5.000000 2.513671 2.513671
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -1.242641 4.028153
       vertex -5.000000 -0.009601 5.000000
       vertex -5.000000 0.009603 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -1.242641 4.028153
       vertex -5.000000 0.009603 5.000000
       vertex -5.000000 0.093323 4.934014
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -1.242641 4.028153
       vertex -5.000000 0.093323 4.934014
       vertex -5.000000 1.686447 3.340894
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 -3.984720 1.866927
       vertex -5.000000 -3.641432 -0.715963
       vertex -5.000000 -5.000000 0.886578
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.984720 1.866927
       vertex -5.000000 -5.000000 0.886578
       vertex -5.000000 -5.000000 1.066713
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.868775 0.994560
       vertex -5.000000 -3.984720 1.866927
       vertex -5.000000 -3.348922 2.368045
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.242642 -3.545530
       vertex -5.000000 0.009602 -5.000000
       vertex -5.000000 -0.009602 -5.000000
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 1.242642 -3.545530
       vertex -5.000000 -0.009602 -5.000000
       vertex -5.000000 -2.296099 -2.302891
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.296098 3.197850
       vertex -5.000000 -1.242641 4.028153
       vertex -5.000000 -0.000000 3.736589
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.538741 1.488606
       vertex -5.000000 3.768895 1.258452
       vertex -5.000000 3.868776 0.994560
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.538741 1.488606
       vertex -5.000000 3.868776 0.994560
       vertex -5.000000 3.736590 -0.000002
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.349002 -3.420070
       vertex -5.000000 1.242642 -3.545530
       vertex -5.000000 1.071068 -3.485281
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.538740 1.866534
       vertex -5.000000 -3.348922 2.368045
       vertex -5.000000 -2.296098 3.197850
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.538740 1.866534
       vertex -5.000000 -2.296098 3.197850
       vertex -5.000000 -0.000000 3.736589
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.538740 1.866534
       vertex -5.000000 -0.000000 3.736589
       vertex -5.000000 1.071070 3.485281
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.768900 1.258436
       vertex -5.000000 -3.538740 1.866534
       vertex -5.000000 -2.956265 2.071071
    endloop
  endfacet
  facet normal -1.000000 -0.000000 -0.000000
    outer loop
       vertex -5.000000 3.538741 1.488606
       vertex -5.000000 3.641432 0.715964
       vertex -5.000000 2.299216 2.299216
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.538741 1.488606
       vertex -5.000000 2.299216 2.299216
       vertex -5.000000 2.513672 2.513672
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.349004 3.420069
       vertex -5.000000 1.686447 3.340894
       vertex -5.000000 2.513671 2.513671
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.349004 3.420069
       vertex -5.000000 2.513671 2.513671
       vertex -5.000000 2.299216 2.299216
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071069
       vertex -5.000000 3.641432 0.715964
       vertex -5.000000 3.736590 -0.000002
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071069
       vertex -5.000000 3.736590 -0.000002
       vertex -5.000000 3.641432 -0.715963
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071069
       vertex -5.000000 3.641432 -0.715963
       vertex -5.000000 3.405246 -0.994564
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.868775 0.994560
       vertex -5.000000 -3.768900 1.258436
       vertex -5.000000 -2.956265 2.071071
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.868775 0.994560
       vertex -5.000000 -2.956265 2.071071
       vertex -5.000000 -2.438503 2.252885
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.868775 0.994560
       vertex -5.000000 -2.438503 2.252885
       vertex -5.000000 -3.405246 -0.994564
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -3.868775 0.994560
       vertex -5.000000 -3.405246 -0.994564
       vertex -5.000000 -3.641432 -0.715963
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.296101 -2.302890
       vertex -5.000000 1.349002 -3.420070
       vertex -5.000000 1.071068 -3.485281
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.296101 -2.302890
       vertex -5.000000 1.071068 -3.485281
       vertex -5.000000 -0.000002 -3.109172
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.000000 3.109172
       vertex -5.000000 1.071070 3.485281
       vertex -5.000000 1.349004 3.420069
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.000000 3.109172
       vertex -5.000000 1.349004 3.420069
       vertex -5.000000 2.296099 2.302891
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 3.109171 0.000001
       vertex -5.000000 3.405246 -0.994564
       vertex -5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 -3.109172 -0.000001
       vertex -5.000000 -2.492628 -2.071070
       vertex -5.000000 -3.405246 -0.994564
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 2.071069
       vertex -5.000000 -2.438503 2.252885
       vertex -5.000000 -2.296100 2.302890
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -5.000000 -5.000000 1.066713
       vertex 0.886578 -5.000000 5.000000
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -5.000000 -5.000000 0.886579
       vertex -5.000000 -5.000000 -5.000000
       vertex -2.642168 -5.000000 -2.642168
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -2.642167 -5.000000 -2.642167
       vertex -5.000000 -5.000000 -5.000000
       vertex -1.066715 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.642167 -5.000000 2.642167
       vertex 5.000000 -5.000000 5.000000
       vertex 1.066713 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 5.000000 -5.000000 -0.886578
       vertex 5.000000 -5.000000 5.000000
       vertex 2.642167 -5.000000 2.642167
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex 5.000000 -5.000000 -0.886578
       vertex 4.407608 -5.000000 -0.000000
       vertex 1.066715 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex 5.000000 -5.000000 -0.886578
       vertex 1.066715 -5.000000 -5.000000
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.689468 -5.000000 4.868296
       vertex 1.226029 -5.000000 4.761567
       vertex 1.066713 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.689468 -5.000000 4.868296
       vertex 1.066713 -5.000000 5.000000
       vertex 0.886578 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.687291 -5.000000 -2.574636
       vertex -2.056763 -5.000000 -3.518288
       vertex -1.066715 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.687291 -5.000000 -2.574636
       vertex -1.066715 -5.000000 -5.000000
       vertex 1.066715 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -3.518287 -5.000000 2.056762
       vertex -5.000000 -5.000000 1.066713
       vertex -5.000000 -5.000000 0.886579
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -3.518287 -5.000000 2.056762
       vertex -5.000000 -5.000000 0.886579
       vertex -2.642168 -5.000000 -2.642168
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -3.518287 -5.000000 2.056762
       vertex -2.642168 -5.000000 -2.642168
       vertex -2.593328 -5.000000 -2.593328
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -2.593327 -5.000000 -2.593327
       vertex -2.642167 -5.000000 -2.642167
       vertex -2.558792 -5.000000 -2.766947
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.305149 -5.000000 3.535535
       vertex -3.518287 -5.000000 2.056762
       vertex -3.237257 -5.000000 0.643930
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex 0.689468 -5.000000 4.868296
       vertex 0.499460 -5.000000 4.741336
       vertex 2.198516 -5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.689468 -5.000000 4.868296
       vertex 2.198516 -5.000000 2.198516
       vertex 2.642167 -5.000000 2.642167
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.689468 -5.000000 4.868296
       vertex 2.642167 -5.000000 2.642167
       vertex 1.226029 -5.000000 4.761567
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 4.037563 -5.000000 -0.553811
       vertex 4.407608 -5.000000 -0.000000
       vertex 2.642167 -5.000000 2.642167
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 4.037563 -5.000000 -0.553811
       vertex 2.642167 -5.000000 2.642167
       vertex 2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.497462 -5.000000 -3.407036
       vertex -2.574636 -5.000000 -2.687291
       vertex -2.558792 -5.000000 -2.766947
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex -1.497462 -5.000000 -3.407036
       vertex -2.558792 -5.000000 -2.766947
       vertex -2.056763 -5.000000 -3.518288
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.833758 -5.000000 2.744416
       vertex 0.499460 -5.000000 4.741336
       vertex -1.305149 -5.000000 3.535535
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.833758 -5.000000 2.744416
       vertex -1.305149 -5.000000 3.535535
       vertex -1.403498 -5.000000 3.388345
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 3.407035 -5.000000 -1.497464
       vertex 4.037563 -5.000000 -0.553811
       vertex 2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex 0.000001 -5.000000 3.109172
       vertex -1.403498 -5.000000 3.388345
       vertex -1.833758 -5.000000 2.744415
    endloop
  endfacet
  facet normal -0.000000 -1.000000 -0.000000
    outer loop
       vertex 1.833758 -5.000000 2.744416
       vertex 1.189827 -5.000000 2.872501
       vertex 2.198516 -5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189828 -5.000000 2.872500
       vertex -1.833758 -5.000000 2.744415
       vertex -2.198516 -5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189827 -5.000000 -2.872501
       vertex 2.687291 -5.000000 -2.574636
       vertex 3.407035 -5.000000 -1.497464
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189827 -5.000000 -2.872501
       vertex 3.407035 -5.000000 -1.497464
       vertex 3.388344 -5.000000 -1.403499
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.744416 -5.000000 -1.833757
       vertex -2.593328 -5.000000 -2.593328
       vertex -2.198516 -5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.497462 -5.000000 -3.407036
       vertex -0.643929 -5.000000 -3.237258
       vertex -2.198516 -5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.497462 -5.000000 -3.407036
       vertex -2.198516 -5.000000 -2.198516
       vertex -2.593327 -5.000000 -2.593327
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.497462 -5.000000 -3.407036
       vertex -2.593327 -5.000000 -2.593327
       vertex -2.574636 -5.000000 -2.687291
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189825 -5.000000 -2.872502
       vertex -0.643929 -5.000000 -3.237258
       vertex -0.000003 -5.000000 -3.109173
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.744417 -5.000000 -1.833757
       vertex 3.388344 -5.000000 -1.403499
       vertex 3.109172 -5.000000 -0.000003
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.872502 -5.000000 1.189825
       vertex -3.237257 -5.000000 0.643930
       vertex -3.109173 -5.000000 0.000003
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.872501 -5.000000 -1.189827
       vertex -2.744416 -5.000000 -1.833757
       vertex -2.198516 -5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.198516 -5.000000 -2.198517
       vertex 2.744417 -5.000000 -1.833757
       vertex 2.872503 -5.000000 -1.189824
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.886578 -5.000000 5.000000
       vertex -5.000000 -0.009600 5.000000
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 2.114610 5.000000 5.000000
       vertex 5.000000 3.072045 5.000000
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -0.326625 5.000000 5.000000
       vertex -5.000000 5.000000 5.000000
       vertex -3.310034 3.310034 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -3.310035 3.310035 5.000000
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 2.352746 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.886577 5.000000 5.000000
       vertex -0.326625 5.000000 5.000000
       vertex -3.310034 3.310034 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.886577 5.000000 5.000000
       vertex -3.310034 3.310034 5.000000
       vertex -2.299216 2.299216 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.299216 2.299216 5.000000
       vertex -3.310035 3.310035 5.000000
       vertex -5.000000 2.352746 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.299216 2.299216 5.000000
       vertex -5.000000 2.352746 5.000000
       vertex -5.000000 0.009603 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 2.800385 -2.800385 5.000000
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -0.009600 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 1.066713 -5.000000 5.000000
       vertex 5.000000 -5.000000 5.000000
       vertex 2.800385 -2.800385 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.420502 4.127431 5.000000
       vertex 2.114610 5.000000 5.000000
       vertex 0.886577 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.420502 4.127431 5.000000
       vertex 0.886577 5.000000 5.000000
       vertex -2.071070 2.492629 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 2.513672 -2.513672 5.000000
       vertex 2.800385 -2.800385 5.000000
       vertex 4.933981 -0.093362 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.421563 -4.605780 5.000000
       vertex 0.886578 -5.000000 5.000000
       vertex 1.066713 -5.000000 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 0.421563 -4.605780 5.000000
       vertex 1.066713 -5.000000 5.000000
       vertex 2.800385 -2.800385 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.421563 -4.605780 5.000000
       vertex 2.800385 -2.800385 5.000000
       vertex 2.513672 -2.513672 5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex -3.545529 -1.242641 5.000000
       vertex -2.302890 2.296101 5.000000
       vertex -5.000000 0.009603 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -3.545529 -1.242641 5.000000
       vertex -5.000000 0.009603 5.000000
       vertex -5.000000 -0.009600 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.420502 4.127431 5.000000
       vertex 1.478137 3.549203 5.000000
       vertex 5.000000 0.027349 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.420502 4.127431 5.000000
       vertex 5.000000 0.027349 5.000000
       vertex 5.000000 3.072045 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -3.109172 -0.000000 5.000000
       vertex -3.545529 -1.242641 5.000000
       vertex -2.302891 -2.296099 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -0.521896 -3.805954 5.000000
       vertex 0.421563 -4.605780 5.000000
       vertex 1.488603 -3.538741 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -0.715963 -3.641432 5.000000
       vertex -0.521896 -3.805954 5.000000
       vertex -0.000002 -3.736589 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.731245 2.296099 5.000000
       vertex 1.478137 3.549203 5.000000
       vertex -0.000001 3.109171 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.607689 1.419657 5.000000
       vertex 2.301125 -2.301125 5.000000
       vertex 2.513672 -2.513672 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.607689 1.419657 5.000000
       vertex 2.513672 -2.513672 5.000000
       vertex 4.933981 -0.093362 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.607689 1.419657 5.000000
       vertex 4.933981 -0.093362 5.000000
       vertex 5.000000 -0.009600 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.607689 1.419657 5.000000
       vertex 5.000000 -0.009600 5.000000
       vertex 5.000000 0.027349 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.071068 -2.956275 5.000000
       vertex 2.513672 -2.513672 5.000000
       vertex 2.301126 -2.301126 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.071067 -2.492629 5.000000
       vertex -0.715963 -3.641432 5.000000
       vertex -0.000002 -3.736589 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.071067 -2.492629 5.000000
       vertex -0.000002 -3.736589 5.000000
       vertex 1.457072 -3.542931 5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex 3.545529 1.242641 5.000000
       vertex 3.607689 1.419657 5.000000
       vertex 2.731245 2.296099 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.545529 1.242641 5.000000
       vertex 2.731245 2.296099 5.000000
       vertex 2.071069 2.492628 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex 2.299216 -2.299216 5.000000
       vertex 2.301125 -2.301125 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.994564 -3.405246 5.000000
       vertex 1.457072 -3.542931 5.000000
       vertex 1.488603 -3.538741 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 0.994564 -3.405246 5.000000
       vertex 1.488603 -3.538741 5.000000
       vertex 2.071068 -2.956275 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.994564 -3.405246 5.000000
       vertex 2.071068 -2.956275 5.000000
       vertex 2.301126 -2.301126 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.994564 -3.405246 5.000000
       vertex 2.301126 -2.301126 5.000000
       vertex 2.299216 -2.299216 5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex 3.485281 1.071070 5.000000
       vertex 3.545529 1.242641 5.000000
       vertex 3.420070 1.349001 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.109172 -0.000001 5.000000
       vertex 3.485281 1.071070 5.000000
       vertex 3.420070 1.349001 5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex 3.109172 -0.000001 5.000000
       vertex 3.420070 1.349001 5.000000
       vertex 2.302889 2.296101 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 0.000002 -3.109172 5.000000
       vertex 0.994564 -3.405246 5.000000
       vertex 2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -5.000000 5.000000 0.886578
       vertex -1.066714 5.000000 -5.000000
       vertex -5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 3.881188 5.000000 -3.881188
       vertex 2.114609 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 2.114609 5.000000 5.000000
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.881189 5.000000 -3.881189
       vertex 5.000000 5.000000 -5.000000
       vertex 4.103734 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -2.198516 5.000000 2.198516
       vertex -5.000000 5.000000 5.000000
       vertex -0.326624 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.037562 5.000000 -0.553812
       vertex -5.000000 5.000000 0.886578
       vertex -5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -4.037562 5.000000 -0.553812
       vertex -5.000000 5.000000 5.000000
       vertex -2.198516 5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.305147 5.000000 3.535537
       vertex -0.326624 5.000000 5.000000
       vertex 0.886576 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.275870 5.000000 4.189286
       vertex -1.350272 5.000000 3.468003
       vertex -1.305147 5.000000 3.535537
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.275870 5.000000 4.189286
       vertex -1.305147 5.000000 3.535537
       vertex 0.886576 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.275870 5.000000 4.189286
       vertex 0.886576 5.000000 5.000000
       vertex 2.114609 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.203803 5.000000 -3.298226
       vertex 4.011130 5.000000 -4.534450
       vertex 4.103734 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.203803 5.000000 -3.298226
       vertex 4.103734 5.000000 -5.000000
       vertex -1.066714 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.558791 5.000000 -2.766949
       vertex -4.037562 5.000000 -0.553812
       vertex -3.237257 5.000000 0.643929
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.275870 5.000000 4.189286
       vertex 2.395921 5.000000 3.585750
       vertex -1.833759 5.000000 2.744414
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.275870 5.000000 4.189286
       vertex -1.833759 5.000000 2.744414
       vertex -1.350272 5.000000 3.468003
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.407035 5.000000 -1.497461
       vertex 3.881188 5.000000 -3.881188
       vertex 2.642167 5.000000 -2.642167
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.686718 5.000000 -4.072097
       vertex 2.642167 5.000000 -2.642167
       vertex 3.881189 5.000000 -3.881189
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 1.686718 5.000000 -4.072097
       vertex 3.881189 5.000000 -3.881189
       vertex 4.011130 5.000000 -4.534450
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 0.147185 5.000000 -3.765865
       vertex -2.203803 5.000000 -3.298226
       vertex -2.558791 5.000000 -2.766949
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 0.147185 5.000000 -3.765865
       vertex -2.558791 5.000000 -2.766949
       vertex -2.744414 5.000000 -1.833759
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.203803 5.000000 -3.298226
       vertex 1.686718 5.000000 -4.072097
       vertex 0.147185 5.000000 -3.765865
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.203803 5.000000 -3.298226
       vertex 0.147185 5.000000 -3.765865
       vertex 0.079656 5.000000 -3.720744
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -3.109172 5.000000 -0.000001
       vertex -3.237257 5.000000 0.643929
       vertex -2.872500 5.000000 1.189827
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.577117 5.000000 -3.422881
       vertex 0.079656 5.000000 -3.720744
       vertex -1.189829 5.000000 -2.872501
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 1.577120 5.000000 3.422880
       vertex 2.395921 5.000000 3.585750
       vertex 2.463384 5.000000 3.246592
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.388345 5.000000 -1.403498
       vertex 3.407035 5.000000 -1.497461
       vertex 2.642167 5.000000 -2.642167
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.388345 5.000000 -1.403498
       vertex 2.642167 5.000000 -2.642167
       vertex 2.198516 5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.203803 5.000000 -3.298226
       vertex 1.577117 5.000000 -3.422881
       vertex 0.643929 5.000000 -3.237258
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.203803 5.000000 -3.298226
       vertex 0.643929 5.000000 -3.237258
       vertex 2.198516 5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex 2.203803 5.000000 -3.298226
       vertex 2.198516 5.000000 -2.198516
       vertex 2.642167 5.000000 -2.642167
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 0.000001 5.000000 3.109172
       vertex 1.577120 5.000000 3.422880
       vertex 2.463384 5.000000 3.246592
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 0.000001 5.000000 3.109172
       vertex 2.463384 5.000000 3.246592
       vertex 2.593327 5.000000 2.593327
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189827 5.000000 2.872500
       vertex -2.198516 5.000000 2.198516
       vertex -1.833759 5.000000 2.744414
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189826 5.000000 2.872501
       vertex 2.593327 5.000000 2.593327
       vertex 2.744414 5.000000 1.833758
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189828 5.000000 -2.872500
       vertex 0.643929 5.000000 -3.237258
       vertex -0.000001 5.000000 -3.109172
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.198517 5.000000 -2.198517
       vertex -2.744414 5.000000 -1.833759
       vertex -2.872500 5.000000 -1.189828
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.198519 5.000000 2.198513
       vertex 2.744414 5.000000 1.833758
       vertex 2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.237258 5.000000 -0.643932
       vertex 3.388345 5.000000 -1.403498
       vertex 2.198516 5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.109172 5.000000 0.000000
       vertex 3.237258 5.000000 -0.643932
       vertex 2.872499 5.000000 -1.189830
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.072044 5.000000
       vertex 5.000000 4.309340 -4.309340
       vertex 5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 3.072044 5.000000
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 4.401134 -5.000000
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 4.309340 -4.309340
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -0.009600 5.000000
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -2.299215 2.299215
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.299215 2.299215
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -5.000000 -0.886579
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -0.009602 -5.000000
       vertex 5.000000 -4.248395 -0.000000
       vertex 5.000000 -5.000000 -0.886579
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -0.009602 -5.000000
       vertex 5.000000 -5.000000 -0.886579
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.805954 -0.521895
       vertex 5.000000 2.299215 -2.299215
       vertex 5.000000 4.309340 -4.309340
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 0.009600 -5.000000
       vertex 5.000000 4.401134 -5.000000
       vertex 5.000000 4.309340 -4.309340
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 0.009600 -5.000000
       vertex 5.000000 4.309340 -4.309340
       vertex 5.000000 2.299215 -2.299215
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.154724 4.377926
       vertex 5.000000 3.072044 5.000000
       vertex 5.000000 0.027348 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.538742 1.488596
       vertex 5.000000 3.154724 4.377926
       vertex 5.000000 0.027348 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 -4.246193 -0.002598
       vertex 5.000000 -4.242640 0.006789
       vertex 5.000000 -4.248395 -0.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.405246 -0.994563
       vertex 5.000000 3.805954 -0.521895
       vertex 5.000000 3.538742 1.488596
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 3.405246 -0.994563
       vertex 5.000000 3.538742 1.488596
       vertex 5.000000 2.296100 2.731241
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -4.246193 -0.002598
       vertex 5.000000 -3.805954 -0.521896
       vertex 5.000000 -3.641431 0.715963
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -4.246193 -0.002598
       vertex 5.000000 -3.641431 0.715963
       vertex 5.000000 -4.242640 0.006789
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -1.349002 -3.420069
       vertex 5.000000 -0.009602 -5.000000
       vertex 5.000000 0.009600 -5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -1.349002 -3.420069
       vertex 5.000000 0.009600 -5.000000
       vertex 5.000000 0.901373 -3.948081
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -1.071071 -3.485281
       vertex 5.000000 0.901373 -3.948081
       vertex 5.000000 2.296100 -2.302889
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 1.045457 3.981888
       vertex 5.000000 0.027348 5.000000
       vertex 5.000000 -0.009600 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 1.045457 3.981888
       vertex 5.000000 -0.009600 5.000000
       vertex 5.000000 -1.349000 3.420069
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.296100 -2.302891
       vertex 5.000000 -1.349002 -3.420069
       vertex 5.000000 -1.071071 -3.485281
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.296100 -2.302891
       vertex 5.000000 -1.071071 -3.485281
       vertex 5.000000 -0.000001 -3.109172
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -3.805954 -0.521896
       vertex 5.000000 -3.405246 -0.994563
       vertex 5.000000 -2.492629 2.071068
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.805954 -0.521896
       vertex 5.000000 -2.492629 2.071068
       vertex 5.000000 -3.641431 0.715963
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.109172 -0.000002
       vertex 5.000000 -3.405246 -0.994563
       vertex 5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492628 2.071070
       vertex 5.000000 2.296100 2.731241
       vertex 5.000000 1.045457 3.981888
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492628 2.071070
       vertex 5.000000 1.045457 3.981888
       vertex 5.000000 0.901375 3.948081
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 1.242642 3.545530
       vertex 5.000000 0.901375 3.948081
       vertex 5.000000 -1.349000 3.420069
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 1.242642 3.545530
       vertex 5.000000 -1.349000 3.420069
       vertex 5.000000 -2.296099 2.302891
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.296101 2.302889
       vertex 5.000000 1.242642 3.545530
       vertex 5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071067
       vertex 5.000000 3.405246 -0.994563
       vertex 5.000000 3.109171 -0.000000
    endloop
  endfacet
  facet normal 0.468140 0.826438 -0.312801
    outer loop
       vertex -2.488345 -5.000000 1.764757
       vertex -2.872500 -5.000000 1.189829
       vertex -3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex -2.488345 -5.000000 1.764757
       vertex -3.919689 -4.242640 1.623588
       vertex -3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.468140 0.826438 -0.312801
    outer loop
       vertex -2.198517 -5.000000 2.198517
       vertex -2.488345 -5.000000 1.764757
       vertex -3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.816085 -0.191480 -0.545290
    outer loop
       vertex -5.000000 0.768310 2.839378
       vertex -5.000000 2.296101 2.302890
       vertex -3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal 0.816084 -0.191480 -0.545290
    outer loop
       vertex -4.242640 0.000000 4.242640
       vertex -5.000000 0.000000 3.109172
       vertex -5.000000 0.768310 2.839378
    endloop
  endfacet
  facet normal 0.816085 -0.191480 -0.545290
    outer loop
       vertex -4.242640 0.000000 4.242640
       vertex -5.000000 0.768310 2.839378
       vertex -3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal 0.552208 -0.826439 0.109841
    outer loop
       vertex -2.974275 5.000000 -0.678174
       vertex -3.919689 4.242640 -1.623588
       vertex -2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.552209 -0.826439 0.109841
    outer loop
       vertex -3.109172 5.000000 0.000000
       vertex -4.242640 4.242640 0.000000
       vertex -3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.552209 -0.826439 0.109841
    outer loop
       vertex -3.109172 5.000000 0.000000
       vertex -3.919689 4.242640 -1.623588
       vertex -2.974275 5.000000 -0.678174
    endloop
  endfacet
  facet normal -0.552209 0.826439 0.109841
    outer loop
       vertex 2.974275 -5.000000 -0.678174
       vertex 3.109172 -5.000000 -0.000000
       vertex 4.242640 -4.242640 -0.000000
    endloop
  endfacet
  facet normal -0.552209 0.826439 0.109841
    outer loop
       vertex 2.974275 -5.000000 -0.678174
       vertex 4.242640 -4.242640 -0.000000
       vertex 3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.552208 0.826439 0.109841
    outer loop
       vertex 2.872500 -5.000000 -1.189829
       vertex 2.974275 -5.000000 -0.678174
       vertex 3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464678
    outer loop
       vertex -5.000000 2.407425 2.171573
       vertex -5.000000 2.492630 2.071068
       vertex -3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464677
    outer loop
       vertex -5.000000 2.407425 2.171573
       vertex -3.919689 4.242640 1.623588
       vertex -3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal 0.695438 -0.548125 -0.464677
    outer loop
       vertex -3.919689 2.296101 3.919689
       vertex -5.000000 2.296101 2.302890
       vertex -5.000000 2.407425 2.171573
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464677
    outer loop
       vertex -3.919689 2.296101 3.919689
       vertex -5.000000 2.407425 2.171573
       vertex -3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex 5.000000 -2.564862 -1.828426
       vertex 4.242640 -4.242640 -0.000000
       vertex 5.000000 -3.109172 -0.000000
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex 5.000000 -2.492630 -2.071068
       vertex 3.919689 -4.242640 -1.623588
       vertex 4.242640 -4.242640 -0.000000
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163172
    outer loop
       vertex 5.000000 -2.492630 -2.071068
       vertex 4.242640 -4.242640 -0.000000
       vertex 5.000000 -2.564862 -1.828426
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex 1.828426 -2.564862 5.000000
       vertex 0.000000 -4.242640 4.242640
       vertex 0.000000 -3.109172 5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex 2.071068 -2.492630 5.000000
       vertex 1.623588 -4.242640 3.919689
       vertex 0.000000 -4.242640 4.242640
    endloop
  endfacet
  facet normal -0.163172 0.548124 -0.820326
    outer loop
       vertex 2.071068 -2.492630 5.000000
       vertex 0.000000 -4.242640 4.242640
       vertex 1.828426 -2.564862 5.000000
    endloop
  endfacet
  facet normal 0.552209 0.826439 0.109841
    outer loop
       vertex -3.007397 -5.000000 -0.511654
       vertex -2.872500 -5.000000 -1.189829
       vertex -3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.552209 0.826439 0.109841
    outer loop
       vertex -3.007397 -5.000000 -0.511654
       vertex -3.919689 -4.242640 -1.623588
       vertex -4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal 0.552209 0.826439 0.109841
    outer loop
       vertex -3.109172 -5.000000 0.000000
       vertex -3.007397 -5.000000 -0.511654
       vertex -4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex -3.919689 -4.242640 1.623588
       vertex -5.000000 -2.492630 2.071068
       vertex -5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex -3.919689 -4.242640 1.623588
       vertex -5.000000 -2.296101 2.302890
       vertex -3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex -3.000000 -4.242640 3.000000
       vertex -3.919689 -4.242640 1.623588
       vertex -3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 -5.000000 2.582671
       vertex -2.198517 -5.000000 2.198517
       vertex -3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 -5.000000 2.582671
       vertex -3.000000 -4.242640 3.000000
       vertex -1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex -1.189829 -5.000000 2.872500
       vertex -1.623588 -5.000000 2.582671
       vertex -1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552208
    outer loop
       vertex -0.678174 5.000000 2.974275
       vertex -1.623588 4.242640 3.919689
       vertex -1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552209
    outer loop
       vertex 0.000000 5.000000 3.109172
       vertex 0.000000 4.242640 4.242640
       vertex -1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552209
    outer loop
       vertex 0.000000 5.000000 3.109172
       vertex -1.623588 4.242640 3.919689
       vertex -0.678174 5.000000 2.974275
    endloop
  endfacet
  facet normal 0.468140 0.826439 0.312801
    outer loop
       vertex -2.582671 -5.000000 -1.623588
       vertex -2.198517 -5.000000 -2.198517
       vertex -3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.468140 0.826439 0.312801
    outer loop
       vertex -2.582671 -5.000000 -1.623588
       vertex -3.000000 -4.242640 -3.000000
       vertex -3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.468140 0.826439 0.312801
    outer loop
       vertex -2.872500 -5.000000 -1.189829
       vertex -2.582671 -5.000000 -1.623588
       vertex -3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.163172 -0.548124 -0.820326
    outer loop
       vertex -1.828426 2.564862 5.000000
       vertex -2.071068 2.492630 5.000000
       vertex -1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex -1.828426 2.564862 5.000000
       vertex -1.623588 4.242640 3.919689
       vertex 0.000000 4.242640 4.242640
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex 0.000000 3.109172 5.000000
       vertex -1.828426 2.564862 5.000000
       vertex 0.000000 4.242640 4.242640
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex 5.000000 -2.947442 0.543277
       vertex 3.919689 -4.242640 1.623588
       vertex 5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex 5.000000 -3.109172 0.000000
       vertex 4.242640 -4.242640 0.000000
       vertex 3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex 5.000000 -3.109172 0.000000
       vertex 3.919689 -4.242640 1.623588
       vertex 5.000000 -2.947442 0.543277
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex 2.582671 5.000000 -1.623588
       vertex 3.919689 4.242640 -1.623588
       vertex 2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex 2.198517 5.000000 -2.198517
       vertex 3.000000 4.242640 -3.000000
       vertex 3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex 2.198517 5.000000 -2.198517
       vertex 3.919689 4.242640 -1.623588
       vertex 2.582671 5.000000 -1.623588
    endloop
  endfacet
  facet normal -0.464678 -0.548124 -0.695439
    outer loop
       vertex 2.171573 2.407425 5.000000
       vertex 2.071068 2.492630 5.000000
       vertex 1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal -0.464677 -0.548124 -0.695439
    outer loop
       vertex 2.171573 2.407425 5.000000
       vertex 1.623588 4.242640 3.919689
       vertex 3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal -0.464677 -0.548125 -0.695438
    outer loop
       vertex 3.919689 2.296101 3.919689
       vertex 2.302890 2.296101 5.000000
       vertex 2.171573 2.407425 5.000000
    endloop
  endfacet
  facet normal -0.464677 -0.548124 -0.695439
    outer loop
       vertex 3.919689 2.296101 3.919689
       vertex 2.171573 2.407425 5.000000
       vertex 3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal -0.163172 -0.548124 0.820326
    outer loop
       vertex 1.828426 2.564862 -5.000000
       vertex 2.071068 2.492630 -5.000000
       vertex 1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex 1.828426 2.564862 -5.000000
       vertex 1.623588 4.242640 -3.919689
       vertex -0.000000 4.242640 -4.242640
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex -0.000000 3.109172 -5.000000
       vertex 1.828426 2.564862 -5.000000
       vertex -0.000000 4.242640 -4.242640
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex 5.000000 1.979074 2.414214
       vertex 4.242640 0.000000 4.242640
       vertex 3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex 5.000000 1.979074 2.414214
       vertex 3.919689 2.296101 3.919689
       vertex 5.000000 2.296101 2.302890
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex 5.000000 0.000000 3.109172
       vertex 4.242640 0.000000 4.242640
       vertex 5.000000 1.979074 2.414214
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex -5.000000 -0.768310 -2.839378
       vertex -3.919689 -2.296101 -3.919689
       vertex -4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal 0.816084 0.191480 0.545290
    outer loop
       vertex -5.000000 -0.768310 -2.839378
       vertex -4.242640 0.000000 -4.242640
       vertex -5.000000 0.000000 -3.109172
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex -5.000000 -2.296101 -2.302890
       vertex -3.919689 -2.296101 -3.919689
       vertex -5.000000 -0.768310 -2.839378
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex 3.919689 -4.242640 -1.623588
       vertex 5.000000 -2.492630 -2.071068
       vertex 5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex 3.919689 -4.242640 -1.623588
       vertex 5.000000 -2.296101 -2.302890
       vertex 3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex 3.000000 -4.242640 -3.000000
       vertex 3.919689 -4.242640 -1.623588
       vertex 3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464678
    outer loop
       vertex 5.000000 2.407425 -2.171573
       vertex 5.000000 2.492630 -2.071068
       vertex 3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464677
    outer loop
       vertex 5.000000 2.407425 -2.171573
       vertex 3.919689 4.242640 -1.623588
       vertex 3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.695438 -0.548125 0.464677
    outer loop
       vertex 3.919689 2.296101 -3.919689
       vertex 5.000000 2.296101 -2.302890
       vertex 5.000000 2.407425 -2.171573
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464677
    outer loop
       vertex 3.919689 2.296101 -3.919689
       vertex 5.000000 2.407425 -2.171573
       vertex 3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552209
    outer loop
       vertex 0.678174 -5.000000 2.974275
       vertex 0.000000 -5.000000 3.109172
       vertex 0.000000 -4.242640 4.242640
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552209
    outer loop
       vertex 0.678174 -5.000000 2.974275
       vertex 0.000000 -4.242640 4.242640
       vertex 1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552208
    outer loop
       vertex 1.189829 -5.000000 2.872500
       vertex 0.678174 -5.000000 2.974275
       vertex 1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.000000 2.947442 0.543277
       vertex -5.000000 3.109172 0.000000
       vertex -4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.000000 2.947442 0.543277
       vertex -4.242640 4.242640 0.000000
       vertex -3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.000000 2.492630 2.071068
       vertex -5.000000 2.947442 0.543277
       vertex -3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal -0.552209 -0.826439 0.109841
    outer loop
       vertex 3.007397 5.000000 -0.511654
       vertex 4.242640 4.242640 -0.000000
       vertex 3.109172 5.000000 -0.000000
    endloop
  endfacet
  facet normal -0.552209 -0.826439 0.109841
    outer loop
       vertex 2.872500 5.000000 -1.189829
       vertex 3.919689 4.242640 -1.623588
       vertex 4.242640 4.242640 -0.000000
    endloop
  endfacet
  facet normal -0.552208 -0.826439 0.109841
    outer loop
       vertex 2.872500 5.000000 -1.189829
       vertex 4.242640 4.242640 -0.000000
       vertex 3.007397 5.000000 -0.511654
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552208
    outer loop
       vertex 0.678174 5.000000 -2.974275
       vertex 1.623588 4.242640 -3.919689
       vertex 1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552209
    outer loop
       vertex -0.000000 5.000000 -3.109172
       vertex -0.000000 4.242640 -4.242640
       vertex 1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552209
    outer loop
       vertex -0.000000 5.000000 -3.109172
       vertex 1.623588 4.242640 -3.919689
       vertex 0.678174 5.000000 -2.974275
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex 5.000000 2.947442 -0.543277
       vertex 5.000000 3.109172 -0.000000
       vertex 4.242640 4.242640 -0.000000
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex 5.000000 2.947442 -0.543277
       vertex 4.242640 4.242640 -0.000000
       vertex 3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex 5.000000 2.492630 -2.071068
       vertex 5.000000 2.947442 -0.543277
       vertex 3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex -3.919689 2.296101 -3.919689
       vertex -3.000000 4.242640 -3.000000
       vertex -3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex -5.000000 2.492630 -2.071068
       vertex -5.000000 2.296101 -2.302890
       vertex -3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex -5.000000 2.492630 -2.071068
       vertex -3.919689 2.296101 -3.919689
       vertex -3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex -5.000000 -1.979074 2.414214
       vertex -5.000000 0.000000 3.109172
       vertex -4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex -3.919689 -2.296101 3.919689
       vertex -5.000000 -2.296101 2.302890
       vertex -5.000000 -1.979074 2.414214
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex -3.919689 -2.296101 3.919689
       vertex -5.000000 -1.979074 2.414214
       vertex -4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal 0.552209 0.826439 -0.109841
    outer loop
       vertex -2.974275 -5.000000 0.678174
       vertex -3.109172 -5.000000 0.000000
       vertex -4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal 0.552209 0.826439 -0.109841
    outer loop
       vertex -2.974275 -5.000000 0.678174
       vertex -4.242640 -4.242640 0.000000
       vertex -3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal 0.552208 0.826439 -0.109841
    outer loop
       vertex -2.872500 -5.000000 1.189829
       vertex -2.974275 -5.000000 0.678174
       vertex -3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal -0.552208 -0.826439 -0.109841
    outer loop
       vertex 2.974275 5.000000 0.678174
       vertex 3.919689 4.242640 1.623588
       vertex 2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal -0.552209 -0.826439 -0.109841
    outer loop
       vertex 3.109172 5.000000 0.000000
       vertex 4.242640 4.242640 0.000000
       vertex 3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal -0.552209 -0.826439 -0.109841
    outer loop
       vertex 3.109172 5.000000 0.000000
       vertex 3.919689 4.242640 1.623588
       vertex 2.974275 5.000000 0.678174
    endloop
  endfacet
  facet normal -0.552209 0.826439 -0.109841
    outer loop
       vertex 3.007397 -5.000000 0.511654
       vertex 2.872500 -5.000000 1.189829
       vertex 3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal -0.552209 0.826439 -0.109841
    outer loop
       vertex 3.007397 -5.000000 0.511654
       vertex 3.919689 -4.242640 1.623588
       vertex 4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal -0.552209 0.826439 -0.109841
    outer loop
       vertex 3.109172 -5.000000 0.000000
       vertex 3.007397 -5.000000 0.511654
       vertex 4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816085
    outer loop
       vertex -2.839378 0.768310 -5.000000
       vertex -2.302890 2.296101 -5.000000
       vertex -3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816084
    outer loop
       vertex -4.242640 0.000000 -4.242640
       vertex -3.109172 0.000000 -5.000000
       vertex -2.839378 0.768310 -5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816085
    outer loop
       vertex -4.242640 0.000000 -4.242640
       vertex -2.839378 0.768310 -5.000000
       vertex -3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552209
    outer loop
       vertex -0.511654 5.000000 -3.007397
       vertex -0.000000 4.242640 -4.242640
       vertex -0.000000 5.000000 -3.109172
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552209
    outer loop
       vertex -1.189829 5.000000 -2.872500
       vertex -1.623588 4.242640 -3.919689
       vertex -0.000000 4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552208
    outer loop
       vertex -1.189829 5.000000 -2.872500
       vertex -0.000000 4.242640 -4.242640
       vertex -0.511654 5.000000 -3.007397
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex -2.839378 -0.768310 5.000000
       vertex -3.919689 -2.296101 3.919689
       vertex -4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816084
    outer loop
       vertex -2.839378 -0.768310 5.000000
       vertex -4.242640 0.000000 4.242640
       vertex -3.109172 0.000000 5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex -2.302890 -2.296101 5.000000
       vertex -3.919689 -2.296101 3.919689
       vertex -2.839378 -0.768310 5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex -2.414214 -1.979074 -5.000000
       vertex -3.109172 0.000000 -5.000000
       vertex -4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex -3.919689 -2.296101 -3.919689
       vertex -2.302890 -2.296101 -5.000000
       vertex -2.414214 -1.979074 -5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex -3.919689 -2.296101 -3.919689
       vertex -2.414214 -1.979074 -5.000000
       vertex -4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal 0.468140 -0.826438 0.312801
    outer loop
       vertex -2.488345 5.000000 -1.764757
       vertex -3.000000 4.242640 -3.000000
       vertex -2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex -2.872500 5.000000 -1.189829
       vertex -3.919689 4.242640 -1.623588
       vertex -3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.468139 -0.826439 0.312801
    outer loop
       vertex -2.872500 5.000000 -1.189829
       vertex -3.000000 4.242640 -3.000000
       vertex -2.488345 5.000000 -1.764757
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex -3.919689 2.296101 3.919689
       vertex -3.000000 4.242640 3.000000
       vertex -1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex -2.071068 2.492630 5.000000
       vertex -2.302890 2.296101 5.000000
       vertex -3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex -2.071068 2.492630 5.000000
       vertex -3.919689 2.296101 3.919689
       vertex -1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex -2.582671 5.000000 1.623588
       vertex -3.919689 4.242640 1.623588
       vertex -2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex -2.198517 5.000000 2.198517
       vertex -3.000000 4.242640 3.000000
       vertex -3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex -2.198517 5.000000 2.198517
       vertex -3.919689 4.242640 1.623588
       vertex -2.582671 5.000000 1.623588
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex 0.543277 2.947442 5.000000
       vertex 0.000000 3.109172 5.000000
       vertex 0.000000 4.242640 4.242640
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex 0.543277 2.947442 5.000000
       vertex 0.000000 4.242640 4.242640
       vertex 1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex 2.071068 2.492630 5.000000
       vertex 0.543277 2.947442 5.000000
       vertex 1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex 1.764757 -5.000000 2.488345
       vertex 1.189829 -5.000000 2.872500
       vertex 1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex 1.764757 -5.000000 2.488345
       vertex 1.623588 -4.242640 3.919689
       vertex 3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal -0.312801 0.826438 -0.468140
    outer loop
       vertex 2.198517 -5.000000 2.198517
       vertex 1.764757 -5.000000 2.488345
       vertex 3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 -0.109841
    outer loop
       vertex -3.007397 5.000000 0.511654
       vertex -4.242640 4.242640 0.000000
       vertex -3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 -0.109841
    outer loop
       vertex -2.872500 5.000000 1.189829
       vertex -3.919689 4.242640 1.623588
       vertex -4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal 0.552208 -0.826439 -0.109841
    outer loop
       vertex -2.872500 5.000000 1.189829
       vertex -4.242640 4.242640 0.000000
       vertex -3.007397 5.000000 0.511654
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex 5.000000 -0.768310 2.839378
       vertex 3.919689 -2.296101 3.919689
       vertex 4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal -0.816084 0.191480 -0.545290
    outer loop
       vertex 5.000000 -0.768310 2.839378
       vertex 4.242640 0.000000 4.242640
       vertex 5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex 5.000000 -2.296101 2.302890
       vertex 3.919689 -2.296101 3.919689
       vertex 5.000000 -0.768310 2.839378
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816085
    outer loop
       vertex 2.839378 0.768310 5.000000
       vertex 2.302890 2.296101 5.000000
       vertex 3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816084
    outer loop
       vertex 4.242640 0.000000 4.242640
       vertex 3.109172 0.000000 5.000000
       vertex 2.839378 0.768310 5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816085
    outer loop
       vertex 4.242640 0.000000 4.242640
       vertex 2.839378 0.768310 5.000000
       vertex 3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex -5.000000 -2.407425 -2.171573
       vertex -3.000000 -4.242640 -3.000000
       vertex -3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.695439 0.548125 0.464677
    outer loop
       vertex -5.000000 -2.407425 -2.171573
       vertex -3.919689 -2.296101 -3.919689
       vertex -5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex -5.000000 -2.492630 -2.071068
       vertex -3.919689 -4.242640 -1.623588
       vertex -3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex -5.000000 -2.492630 -2.071068
       vertex -3.000000 -4.242640 -3.000000
       vertex -5.000000 -2.407425 -2.171573
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex -0.543277 2.947442 -5.000000
       vertex -0.000000 3.109172 -5.000000
       vertex -0.000000 4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex -0.543277 2.947442 -5.000000
       vertex -0.000000 4.242640 -4.242640
       vertex -1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex -2.071068 2.492630 -5.000000
       vertex -0.543277 2.947442 -5.000000
       vertex -1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552209
    outer loop
       vertex 0.511654 5.000000 3.007397
       vertex 0.000000 4.242640 4.242640
       vertex 0.000000 5.000000 3.109172
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552209
    outer loop
       vertex 1.189829 5.000000 2.872500
       vertex 1.623588 4.242640 3.919689
       vertex 0.000000 4.242640 4.242640
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552208
    outer loop
       vertex 1.189829 5.000000 2.872500
       vertex 0.000000 4.242640 4.242640
       vertex 0.511654 5.000000 3.007397
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex 2.171573 -2.407425 -5.000000
       vertex 3.000000 -4.242640 -3.000000
       vertex 3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.464677 0.548125 0.695439
    outer loop
       vertex 2.171573 -2.407425 -5.000000
       vertex 3.919689 -2.296101 -3.919689
       vertex 2.302890 -2.296101 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex 2.071068 -2.492630 -5.000000
       vertex 1.623588 -4.242640 -3.919689
       vertex 3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex 2.071068 -2.492630 -5.000000
       vertex 3.000000 -4.242640 -3.000000
       vertex 2.171573 -2.407425 -5.000000
    endloop
  endfacet
  facet normal -0.816085 -0.191480 0.545290
    outer loop
       vertex 5.000000 0.768310 -2.839378
       vertex 5.000000 2.296101 -2.302890
       vertex 3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.816084 -0.191480 0.545290
    outer loop
       vertex 4.242640 0.000000 -4.242640
       vertex 5.000000 0.000000 -3.109172
       vertex 5.000000 0.768310 -2.839378
    endloop
  endfacet
  facet normal -0.816085 -0.191480 0.545290
    outer loop
       vertex 4.242640 0.000000 -4.242640
       vertex 5.000000 0.768310 -2.839378
       vertex 3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163172
    outer loop
       vertex 5.000000 2.564862 1.828426
       vertex 5.000000 2.492630 2.071068
       vertex 3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex 5.000000 2.564862 1.828426
       vertex 3.919689 4.242640 1.623588
       vertex 4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex 5.000000 3.109172 0.000000
       vertex 5.000000 2.564862 1.828426
       vertex 4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex -1.828426 -2.564862 -5.000000
       vertex -0.000000 -4.242640 -4.242640
       vertex -0.000000 -3.109172 -5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex -2.071068 -2.492630 -5.000000
       vertex -1.623588 -4.242640 -3.919689
       vertex -0.000000 -4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.163172 0.548124 0.820326
    outer loop
       vertex -2.071068 -2.492630 -5.000000
       vertex -0.000000 -4.242640 -4.242640
       vertex -1.828426 -2.564862 -5.000000
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex 3.919689 2.296101 -3.919689
       vertex 3.000000 4.242640 -3.000000
       vertex 1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex 2.071068 2.492630 -5.000000
       vertex 2.302890 2.296101 -5.000000
       vertex 3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex 2.071068 2.492630 -5.000000
       vertex 3.919689 2.296101 -3.919689
       vertex 1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex -1.623588 5.000000 -2.582671
       vertex -1.623588 4.242640 -3.919689
       vertex -1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex -2.198517 5.000000 -2.198517
       vertex -3.000000 4.242640 -3.000000
       vertex -1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex -2.198517 5.000000 -2.198517
       vertex -1.623588 4.242640 -3.919689
       vertex -1.623588 5.000000 -2.582671
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163172
    outer loop
       vertex -5.000000 2.564862 -1.828426
       vertex -5.000000 2.492630 -2.071068
       vertex -3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex -5.000000 2.564862 -1.828426
       vertex -3.919689 4.242640 -1.623588
       vertex -4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex -5.000000 3.109172 0.000000
       vertex -5.000000 2.564862 -1.828426
       vertex -4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex 2.414214 1.979074 -5.000000
       vertex 4.242640 0.000000 -4.242640
       vertex 3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex 2.414214 1.979074 -5.000000
       vertex 3.919689 2.296101 -3.919689
       vertex 2.302890 2.296101 -5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex 3.109172 0.000000 -5.000000
       vertex 4.242640 0.000000 -4.242640
       vertex 2.414214 1.979074 -5.000000
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 -5.000000 -2.582671
       vertex 2.198517 -5.000000 -2.198517
       vertex 3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 -5.000000 -2.582671
       vertex 3.000000 -4.242640 -3.000000
       vertex 1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex 1.189829 -5.000000 -2.872500
       vertex 1.623588 -5.000000 -2.582671
       vertex 1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex 5.000000 -2.407425 2.171573
       vertex 3.000000 -4.242640 3.000000
       vertex 3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal -0.695439 0.548125 -0.464677
    outer loop
       vertex 5.000000 -2.407425 2.171573
       vertex 3.919689 -2.296101 3.919689
       vertex 5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex 5.000000 -2.492630 2.071068
       vertex 3.919689 -4.242640 1.623588
       vertex 3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex 5.000000 -2.492630 2.071068
       vertex 3.000000 -4.242640 3.000000
       vertex 5.000000 -2.407425 2.171573
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex -2.414214 1.979074 5.000000
       vertex -4.242640 0.000000 4.242640
       vertex -3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex -2.414214 1.979074 5.000000
       vertex -3.919689 2.296101 3.919689
       vertex -2.302890 2.296101 5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex -3.109172 0.000000 5.000000
       vertex -4.242640 0.000000 4.242640
       vertex -2.414214 1.979074 5.000000
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex -5.000000 1.979074 -2.414214
       vertex -4.242640 0.000000 -4.242640
       vertex -3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex -5.000000 1.979074 -2.414214
       vertex -3.919689 2.296101 -3.919689
       vertex -5.000000 2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex -5.000000 0.000000 -3.109172
       vertex -4.242640 0.000000 -4.242640
       vertex -5.000000 1.979074 -2.414214
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex -2.171573 -2.407425 5.000000
       vertex -3.000000 -4.242640 3.000000
       vertex -3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal 0.464677 0.548125 -0.695439
    outer loop
       vertex -2.171573 -2.407425 5.000000
       vertex -3.919689 -2.296101 3.919689
       vertex -2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex -2.071068 -2.492630 5.000000
       vertex -1.623588 -4.242640 3.919689
       vertex -3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex -2.071068 -2.492630 5.000000
       vertex -3.000000 -4.242640 3.000000
       vertex -2.171573 -2.407425 5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex 0.543277 -2.947442 -5.000000
       vertex 1.623588 -4.242640 -3.919689
       vertex 2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex -0.000000 -3.109172 -5.000000
       vertex -0.000000 -4.242640 -4.242640
       vertex 1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex -0.000000 -3.109172 -5.000000
       vertex 1.623588 -4.242640 -3.919689
       vertex 0.543277 -2.947442 -5.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex -5.000000 -2.564862 1.828426
       vertex -4.242640 -4.242640 0.000000
       vertex -5.000000 -3.109172 0.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex -5.000000 -2.492630 2.071068
       vertex -3.919689 -4.242640 1.623588
       vertex -4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163172
    outer loop
       vertex -5.000000 -2.492630 2.071068
       vertex -4.242640 -4.242640 0.000000
       vertex -5.000000 -2.564862 1.828426
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex 2.414214 -1.979074 5.000000
       vertex 3.109172 0.000000 5.000000
       vertex 4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex 3.919689 -2.296101 3.919689
       vertex 2.302890 -2.296101 5.000000
       vertex 2.414214 -1.979074 5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex 3.919689 -2.296101 3.919689
       vertex 2.414214 -1.979074 5.000000
       vertex 4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex -1.764757 -5.000000 -2.488345
       vertex -1.189829 -5.000000 -2.872500
       vertex -1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex -1.764757 -5.000000 -2.488345
       vertex -1.623588 -4.242640 -3.919689
       vertex -3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.312801 0.826438 0.468140
    outer loop
       vertex -2.198517 -5.000000 -2.198517
       vertex -1.764757 -5.000000 -2.488345
       vertex -3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.312801 -0.826438 0.468140
    outer loop
       vertex 1.764757 5.000000 -2.488345
       vertex 3.000000 4.242640 -3.000000
       vertex 2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex 1.189829 5.000000 -2.872500
       vertex 1.623588 4.242640 -3.919689
       vertex 3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468139
    outer loop
       vertex 1.189829 5.000000 -2.872500
       vertex 3.000000 4.242640 -3.000000
       vertex 1.764757 5.000000 -2.488345
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex 1.623588 5.000000 2.582671
       vertex 1.623588 4.242640 3.919689
       vertex 1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex 2.198517 5.000000 2.198517
       vertex 3.000000 4.242640 3.000000
       vertex 1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex 2.198517 5.000000 2.198517
       vertex 1.623588 4.242640 3.919689
       vertex 1.623588 5.000000 2.582671
    endloop
  endfacet
  facet normal -0.468140 0.826439 -0.312801
    outer loop
       vertex 2.582671 -5.000000 1.623588
       vertex 2.198517 -5.000000 2.198517
       vertex 3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal -0.468140 0.826439 -0.312801
    outer loop
       vertex 2.582671 -5.000000 1.623588
       vertex 3.000000 -4.242640 3.000000
       vertex 3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal -0.468140 0.826439 -0.312801
    outer loop
       vertex 2.872500 -5.000000 1.189829
       vertex 2.582671 -5.000000 1.623588
       vertex 3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal 0.312801 -0.826438 -0.468140
    outer loop
       vertex -1.764757 5.000000 2.488345
       vertex -3.000000 4.242640 3.000000
       vertex -2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex -1.623588 4.242640 3.919689
       vertex -3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468139
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex -3.000000 4.242640 3.000000
       vertex -1.764757 5.000000 2.488345
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex 2.839378 -0.768310 -5.000000
       vertex 3.919689 -2.296101 -3.919689
       vertex 4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816084
    outer loop
       vertex 2.839378 -0.768310 -5.000000
       vertex 4.242640 0.000000 -4.242640
       vertex 3.109172 0.000000 -5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex 2.302890 -2.296101 -5.000000
       vertex 3.919689 -2.296101 -3.919689
       vertex 2.839378 -0.768310 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex 1.623588 -4.242640 3.919689
       vertex 2.071068 -2.492630 5.000000
       vertex 2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex 1.623588 -4.242640 3.919689
       vertex 2.302890 -2.296101 5.000000
       vertex 3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex 3.000000 -4.242640 3.000000
       vertex 1.623588 -4.242640 3.919689
       vertex 3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552209
    outer loop
       vertex 0.511654 -5.000000 -3.007397
       vertex 1.189829 -5.000000 -2.872500
       vertex 1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552209
    outer loop
       vertex 0.511654 -5.000000 -3.007397
       vertex 1.623588 -4.242640 -3.919689
       vertex -0.000000 -4.242640 -4.242640
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552209
    outer loop
       vertex -0.000000 -5.000000 -3.109172
       vertex 0.511654 -5.000000 -3.007397
       vertex -0.000000 -4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552209
    outer loop
       vertex -0.678174 -5.000000 -2.974275
       vertex -0.000000 -5.000000 -3.109172
       vertex -0.000000 -4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552209
    outer loop
       vertex -0.678174 -5.000000 -2.974275
       vertex -0.000000 -4.242640 -4.242640
       vertex -1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552208
    outer loop
       vertex -1.189829 -5.000000 -2.872500
       vertex -0.678174 -5.000000 -2.974275
       vertex -1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.468140 0.826438 0.312801
    outer loop
       vertex 2.488345 -5.000000 -1.764757
       vertex 2.872500 -5.000000 -1.189829
       vertex 3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex 2.488345 -5.000000 -1.764757
       vertex 3.919689 -4.242640 -1.623588
       vertex 3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.468140 0.826438 0.312801
    outer loop
       vertex 2.198517 -5.000000 -2.198517
       vertex 2.488345 -5.000000 -1.764757
       vertex 3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex -5.000000 -2.947442 -0.543277
       vertex -3.919689 -4.242640 -1.623588
       vertex -5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex -5.000000 -3.109172 0.000000
       vertex -4.242640 -4.242640 0.000000
       vertex -3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex -5.000000 -3.109172 0.000000
       vertex -3.919689 -4.242640 -1.623588
       vertex -5.000000 -2.947442 -0.543277
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex -1.623588 -4.242640 -3.919689
       vertex -2.071068 -2.492630 -5.000000
       vertex -2.302890 -2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex -1.623588 -4.242640 -3.919689
       vertex -2.302890 -2.296101 -5.000000
       vertex -3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex -3.000000 -4.242640 -3.000000
       vertex -1.623588 -4.242640 -3.919689
       vertex -3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex -0.543277 -2.947442 5.000000
       vertex -1.623588 -4.242640 3.919689
       vertex -2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex 0.000000 -3.109172 5.000000
       vertex 0.000000 -4.242640 4.242640
       vertex -1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex 0.000000 -3.109172 5.000000
       vertex -1.623588 -4.242640 3.919689
       vertex -0.543277 -2.947442 5.000000
    endloop
  endfacet
  facet normal 0.464678 -0.548124 0.695439
    outer loop
       vertex -2.171573 2.407425 -5.000000
       vertex -2.071068 2.492630 -5.000000
       vertex -1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.464677 -0.548124 0.695439
    outer loop
       vertex -2.171573 2.407425 -5.000000
       vertex -1.623588 4.242640 -3.919689
       vertex -3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548125 0.695438
    outer loop
       vertex -3.919689 2.296101 -3.919689
       vertex -2.302890 2.296101 -5.000000
       vertex -2.171573 2.407425 -5.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 0.695439
    outer loop
       vertex -3.919689 2.296101 -3.919689
       vertex -2.171573 2.407425 -5.000000
       vertex -3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex 5.000000 -1.979074 -2.414214
       vertex 5.000000 0.000000 -3.109172
       vertex 4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex 3.919689 -2.296101 -3.919689
       vertex 5.000000 -2.296101 -2.302890
       vertex 5.000000 -1.979074 -2.414214
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex 3.919689 -2.296101 -3.919689
       vertex 5.000000 -1.979074 -2.414214
       vertex 4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal -0.468140 -0.826438 -0.312801
    outer loop
       vertex 2.488345 5.000000 1.764757
       vertex 3.000000 4.242640 3.000000
       vertex 2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex 2.872500 5.000000 1.189829
       vertex 3.919689 4.242640 1.623588
       vertex 3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal -0.468139 -0.826439 -0.312801
    outer loop
       vertex 2.872500 5.000000 1.189829
       vertex 3.000000 4.242640 3.000000
       vertex 2.488345 5.000000 1.764757
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex 3.919689 2.296101 3.919689
       vertex 3.000000 4.242640 3.000000
       vertex 3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex 5.000000 2.492630 2.071068
       vertex 5.000000 2.296101 2.302890
       vertex 3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex 5.000000 2.492630 2.071068
       vertex 3.919689 2.296101 3.919689
       vertex 3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552209
    outer loop
       vertex -0.511654 -5.000000 3.007397
       vertex -1.189829 -5.000000 2.872500
       vertex -1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552209
    outer loop
       vertex -0.511654 -5.000000 3.007397
       vertex -1.623588 -4.242640 3.919689
       vertex 0.000000 -4.242640 4.242640
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552209
    outer loop
       vertex 0.000000 -5.000000 3.109172
       vertex -0.511654 -5.000000 3.007397
       vertex 0.000000 -4.242640 4.242640
    endloop
  endfacet
endsolid
```

### Union

```stl
solid ScadSharp
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 3.298226 -3.298226 5.000000
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -3.072045 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 3.298226 -3.298226 5.000000
       vertex -5.000000 -4.401134 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex -5.000000 -4.401134 5.000000
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -5.000000 -1.620707 5.000000
       vertex -5.000000 -4.401134 5.000000
       vertex -3.293709 -4.174352 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -4.791841 4.791841 5.000000
       vertex -4.864937 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -4.791840 4.791840 5.000000
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 -1.620707 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -4.791840 4.791840 5.000000
       vertex -5.000000 -1.620707 5.000000
       vertex -3.293709 -4.174352 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -4.791840 4.791840 5.000000
       vertex -3.293709 -4.174352 5.000000
       vertex -1.716927 -3.964784 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.301126 -2.301126 5.000000
       vertex 3.298226 -3.298226 5.000000
       vertex 5.000000 -3.072045 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 2.301126 -2.301126 5.000000
       vertex 5.000000 -3.072045 5.000000
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.301126 -2.301126 5.000000
       vertex 5.000000 5.000000 5.000000
       vertex 4.864937 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 3.298226 -3.298226 5.000000
       vertex 2.301126 -2.301126 5.000000
       vertex 1.885035 -3.486052 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -4.791841 4.791841 5.000000
       vertex -4.174983 4.174983 5.000000
       vertex -4.010877 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex -4.791841 4.791841 5.000000
       vertex -4.010877 5.000000 5.000000
       vertex -4.864937 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -4.174983 4.174983 5.000000
       vertex -4.791840 4.791840 5.000000
       vertex -4.319718 3.447349 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -4.174983 4.174983 5.000000
       vertex -3.298226 3.298226 5.000000
       vertex 4.637737 4.352987 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -4.174983 4.174983 5.000000
       vertex 4.637737 4.352987 5.000000
       vertex 4.864937 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -4.174983 4.174983 5.000000
       vertex 4.864937 5.000000 5.000000
       vertex -4.010877 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -3.298225 3.298225 5.000000
       vertex -4.174983 4.174983 5.000000
       vertex -4.319718 3.447349 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -3.298225 3.298225 5.000000
       vertex -4.319718 3.447349 5.000000
       vertex -4.224140 3.175164 5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex 3.109172 -0.000000 5.000000
       vertex 4.637737 4.352987 5.000000
       vertex 1.716927 3.964785 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -3.298226 3.298226 5.000000
       vertex -2.299216 2.299216 5.000000
       vertex -0.715965 3.641432 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.299217 2.299217 5.000000
       vertex -3.298225 3.298225 5.000000
       vertex -4.224140 3.175164 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.299217 2.299217 5.000000
       vertex -4.224140 3.175164 5.000000
       vertex -3.545530 1.242643 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302891 -2.296099 5.000000
       vertex -1.716927 -3.964784 5.000000
       vertex -0.521894 -3.805954 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.252885 2.438504 5.000000
       vertex 1.716927 3.964785 5.000000
       vertex -0.715965 3.641432 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.252885 2.438504 5.000000
       vertex -0.715965 3.641432 5.000000
       vertex -0.994564 3.405246 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302891 2.296099 5.000000
       vertex 2.252885 2.438504 5.000000
       vertex 2.071067 2.492630 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302889 2.296104 5.000000
       vertex -3.545530 1.242643 5.000000
       vertex -3.109172 -0.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -0.994562 -3.405246 5.000000
       vertex -0.521894 -3.805954 5.000000
       vertex 1.885035 -3.486052 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -0.994562 -3.405246 5.000000
       vertex 1.885035 -3.486052 5.000000
       vertex 2.252884 -2.438504 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.071067 -2.492629 5.000000
       vertex -0.994562 -3.405246 5.000000
       vertex 0.000000 -3.109172 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 0.000000 3.109172 5.000000
       vertex -0.994564 3.405246 5.000000
       vertex -2.071070 2.492629 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex 2.299216 -2.299216 5.000000
       vertex 2.301126 -2.301126 5.000000
       vertex 2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.301126 -2.301126 5.000000
       vertex 2.299216 -2.299216 5.000000
       vertex 2.071069 -2.492630 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.301126 -2.301126 5.000000
       vertex 2.071069 -2.492630 5.000000
       vertex 2.252884 -2.438504 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 4.791841 -4.791841
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 5.000000 -4.864936
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 4.791841 -4.791841
       vertex 5.000000 -5.000000 -1.353409
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 -5.000000 -1.353409
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -3.844275 3.844275
       vertex 5.000000 -3.072045 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.844275 3.844275
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -5.000000 2.114612
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 4.634643 -4.634643
       vertex 5.000000 4.791841 -4.791841
       vertex 5.000000 5.000000 -4.864936
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 4.634643 -4.634643
       vertex 5.000000 5.000000 -4.864936
       vertex 5.000000 5.000000 -1.392058
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 4.791841 -4.791841
       vertex 5.000000 4.634642 -4.634642
       vertex 5.000000 4.623587 -4.732758
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 3.310035 -3.310035
       vertex 5.000000 4.634643 -4.634643
       vertex 5.000000 5.000000 -1.392058
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.310035 -3.310035
       vertex 5.000000 5.000000 -1.392058
       vertex 5.000000 5.000000 -0.326625
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 4.634642 -4.634642
       vertex 5.000000 3.310034 -3.310034
       vertex 5.000000 2.855767 -4.111983
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 4.634642 -4.634642
       vertex 5.000000 2.855767 -4.111983
       vertex 5.000000 4.623587 -4.732758
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -3.844275 3.844275
       vertex 5.000000 -3.026478 3.026478
       vertex 5.000000 5.000000 4.909754
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.844275 3.844275
       vertex 5.000000 5.000000 4.909754
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.844275 3.844275
       vertex 5.000000 5.000000 5.000000
       vertex 5.000000 -3.072045 5.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.026478 3.026478
       vertex 5.000000 -3.844275 3.844275
       vertex 5.000000 -4.644356 2.646871
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 5.000000 4.103734
       vertex 5.000000 5.000000 4.909754
       vertex 5.000000 4.534426 4.800515
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 5.000000 2.563426
       vertex 5.000000 5.000000 4.103734
       vertex 5.000000 4.534426 4.800515
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 5.000000 2.563426
       vertex 5.000000 4.534426 4.800515
       vertex 5.000000 0.000000 3.736589
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.945911 2.810749
       vertex 5.000000 -4.644356 2.646871
       vertex 5.000000 -5.000000 2.114612
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.945911 2.810749
       vertex 5.000000 -5.000000 2.114612
       vertex 5.000000 -5.000000 -1.353409
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.945911 2.810749
       vertex 5.000000 -5.000000 -1.353409
       vertex 5.000000 -2.438503 -2.252886
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 3.026478 -3.026478
       vertex 5.000000 3.310035 -3.310035
       vertex 5.000000 3.538742 -2.906284
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.310034 -3.310034
       vertex 5.000000 3.026478 -3.026478
       vertex 5.000000 1.071069 -3.485281
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.310034 -3.310034
       vertex 5.000000 1.071069 -3.485281
       vertex 5.000000 2.855767 -4.111983
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.344253 -2.951917
       vertex 5.000000 3.538742 -2.906284
       vertex 5.000000 5.000000 -0.326625
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.344253 -2.951917
       vertex 5.000000 5.000000 -0.326625
       vertex 5.000000 5.000000 2.563426
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.344253 -2.951917
       vertex 5.000000 5.000000 2.563426
       vertex 5.000000 4.105196 2.773376
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -3.026478 3.026478
       vertex 5.000000 -2.301126 2.301126
       vertex 5.000000 1.071070 3.485281
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.026478 3.026478
       vertex 5.000000 1.071070 3.485281
       vertex 5.000000 0.000000 3.736589
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.301125 2.301125
       vertex 5.000000 -3.026478 3.026478
       vertex 5.000000 -3.945911 2.810749
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.301125 2.301125
       vertex 5.000000 -3.945911 2.810749
       vertex 5.000000 -3.652892 1.826449
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -3.538740 1.866534
       vertex 5.000000 -3.652892 1.826449
       vertex 5.000000 -3.609316 1.680070
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.956275 2.071068
       vertex 5.000000 -3.538740 1.866534
       vertex 5.000000 -3.609316 1.680070
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.956275 2.071068
       vertex 5.000000 -3.609316 1.680070
       vertex 5.000000 -3.549202 1.478138
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.438504 2.252884
       vertex 5.000000 -2.956275 2.071068
       vertex 5.000000 -3.549202 1.478138
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.438504 2.252884
       vertex 5.000000 -3.549202 1.478138
       vertex 5.000000 -3.109171 0.000001
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 3.984721 1.866927
       vertex 5.000000 4.105196 2.773376
       vertex 5.000000 2.296096 3.197850
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -2.301126 2.301126
       vertex 5.000000 -2.299216 2.299216
       vertex 5.000000 -2.296100 2.302891
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.299216 2.299216
       vertex 5.000000 -2.301125 2.301125
       vertex 5.000000 -2.438504 2.252884
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.299216 2.299216
       vertex 5.000000 -2.438504 2.252884
       vertex 5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.301125 -2.301125
       vertex 5.000000 3.026478 -3.026478
       vertex 5.000000 3.344253 -2.951917
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 2.301125 -2.301125
       vertex 5.000000 3.344253 -2.951917
       vertex 5.000000 3.486051 -1.885036
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.026478 -3.026478
       vertex 5.000000 2.301126 -2.301126
       vertex 5.000000 -0.000000 -3.109173
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.026478 -3.026478
       vertex 5.000000 -0.000000 -3.109173
       vertex 5.000000 1.071069 -3.485281
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 3.805954 0.521896
       vertex 5.000000 3.984721 1.866927
       vertex 5.000000 2.296096 3.197850
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 3.805954 0.521896
       vertex 5.000000 2.296096 3.197850
       vertex 5.000000 1.349001 3.420070
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 3.542931 -1.457072
       vertex 5.000000 3.805954 0.521896
       vertex 5.000000 2.492629 2.071068
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.296100 2.302890
       vertex 5.000000 1.349001 3.420070
       vertex 5.000000 1.071070 3.485281
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.296100 2.302890
       vertex 5.000000 1.071070 3.485281
       vertex 5.000000 -0.000002 3.109172
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.438503 -2.252885
       vertex 5.000000 3.486051 -1.885036
       vertex 5.000000 3.542931 -1.457072
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.438503 -2.252885
       vertex 5.000000 3.542931 -1.457072
       vertex 5.000000 3.109171 -0.000000
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492628 -2.071071
       vertex 5.000000 -2.438503 -2.252886
       vertex 5.000000 -2.296100 -2.302891
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 2.299216 -2.299216
       vertex 5.000000 2.301125 -2.301125
       vertex 5.000000 2.438503 -2.252885
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.299216 -2.299216
       vertex 5.000000 2.438503 -2.252885
       vertex 5.000000 2.492628 -2.071071
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.301126 -2.301126
       vertex 5.000000 2.299216 -2.299216
       vertex 5.000000 2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 4.945901 5.000000 -4.945901
       vertex 5.000000 5.000000 -4.864937
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 4.945902 5.000000 -4.945902
       vertex 5.000000 5.000000 -5.000000
       vertex 4.909753 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex 4.945901 5.000000 -4.945901
       vertex 3.554861 5.000000 -3.554861
       vertex 5.000000 5.000000 -1.392060
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 4.945901 5.000000 -4.945901
       vertex 5.000000 5.000000 -1.392060
       vertex 5.000000 5.000000 -4.864937
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.554861 5.000000 -3.554861
       vertex 4.945902 5.000000 -4.945902
       vertex 4.909753 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.554861 5.000000 -3.554861
       vertex 4.909753 5.000000 -5.000000
       vertex 2.589254 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.010875 5.000000 -5.000000
       vertex -5.000000 5.000000 -5.000000
       vertex -5.000000 5.000000 -0.027335
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex 3.554861 5.000000 -3.554861
       vertex 2.198517 5.000000 -2.198517
       vertex 5.000000 5.000000 -0.326626
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.554861 5.000000 -3.554861
       vertex 5.000000 5.000000 -0.326626
       vertex 5.000000 5.000000 -1.392060
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.198517 5.000000 -2.198517
       vertex 3.554861 5.000000 -3.554861
       vertex 2.589254 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.198517 5.000000 -2.198517
       vertex 2.589254 5.000000 -5.000000
       vertex -1.994200 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -0.553811 5.000000 -4.037563
       vertex -1.994200 5.000000 -5.000000
       vertex -4.010875 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -0.553811 5.000000 -4.037563
       vertex -4.010875 5.000000 -5.000000
       vertex -4.761564 5.000000 -1.226034
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -4.945901 5.000000 4.945901
       vertex -5.000000 5.000000 5.000000
       vertex -4.864936 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -4.945901 5.000000 4.945901
       vertex -5.000000 5.000000 4.909753
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 4.864936 5.000000 5.000000
       vertex 5.000000 5.000000 5.000000
       vertex 5.000000 5.000000 4.909754
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.174984 5.000000 4.174984
       vertex -4.945901 5.000000 4.945901
       vertex -4.864936 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -4.174984 5.000000 4.174984
       vertex -4.864936 5.000000 5.000000
       vertex -4.010877 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.945901 5.000000 4.945901
       vertex -4.174984 5.000000 4.174984
       vertex -5.000000 5.000000 0.027349
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.945901 5.000000 4.945901
       vertex -5.000000 5.000000 0.027349
       vertex -5.000000 5.000000 4.909753
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189831 5.000000 -2.872499
       vertex -0.553811 5.000000 -4.037563
       vertex -1.497463 5.000000 -3.407035
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex -4.174984 5.000000 4.174984
       vertex -4.010877 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex -4.010877 5.000000 5.000000
       vertex 4.864936 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex 4.864936 5.000000 5.000000
       vertex 5.000000 5.000000 4.909754
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex 5.000000 5.000000 4.909754
       vertex 5.000000 5.000000 4.103734
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.174984 5.000000 4.174984
       vertex -2.593327 5.000000 2.593327
       vertex -4.567714 5.000000 2.200597
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 3.223589 5.000000 3.750384
       vertex 5.000000 5.000000 4.103734
       vertex 5.000000 5.000000 2.563426
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.761567 5.000000 1.226030
       vertex -4.567714 5.000000 2.200597
       vertex -2.766950 5.000000 2.558791
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.781076 5.000000 3.377884
       vertex 5.000000 5.000000 2.563426
       vertex 5.000000 5.000000 -0.326626
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 3.781076 5.000000 3.377884
       vertex 5.000000 5.000000 -0.326626
       vertex 2.744415 5.000000 -1.833760
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 1.577118 5.000000 3.422881
       vertex 3.223589 5.000000 3.750384
       vertex 3.781076 5.000000 3.377884
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.577118 5.000000 3.422881
       vertex 3.781076 5.000000 3.377884
       vertex 3.705803 5.000000 2.999460
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.868294 5.000000 -0.689465
       vertex -5.000000 5.000000 -0.027335
       vertex -5.000000 5.000000 0.027349
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -4.868294 5.000000 -0.689465
       vertex -5.000000 5.000000 0.027349
       vertex -4.761567 5.000000 1.226030
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.868294 5.000000 -0.689465
       vertex -4.761567 5.000000 1.226030
       vertex -2.766950 5.000000 2.558791
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -4.868294 5.000000 -0.689465
       vertex -2.766950 5.000000 2.558791
       vertex -2.687290 5.000000 2.574636
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.574635 5.000000 -2.687292
       vertex -4.761564 5.000000 -1.226034
       vertex -4.868294 5.000000 -0.689465
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.574635 5.000000 -2.687292
       vertex -4.868294 5.000000 -0.689465
       vertex -3.407036 5.000000 1.497462
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -0.000003 5.000000 -3.109172
       vertex -1.497463 5.000000 -3.407035
       vertex -2.574635 5.000000 -2.687292
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -0.000003 5.000000 -3.109172
       vertex -2.574635 5.000000 -2.687292
       vertex -2.593326 5.000000 -2.593327
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 1.497461 5.000000 3.407036
       vertex 1.577118 5.000000 3.422881
       vertex 3.705803 5.000000 2.999460
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.497461 5.000000 3.407036
       vertex 3.705803 5.000000 2.999460
       vertex 3.518288 5.000000 2.056762
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex -3.109172 5.000000 0.000001
       vertex -3.407036 5.000000 1.497462
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex -3.407036 5.000000 1.497462
       vertex -2.687290 5.000000 2.574636
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.642168 5.000000 2.642167
       vertex 3.518288 5.000000 2.056762
       vertex 3.407036 5.000000 1.497460
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex 0.000001 5.000000 3.109172
       vertex 1.497461 5.000000 3.407036
       vertex 2.642168 5.000000 2.642167
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 0.000001 5.000000 3.109172
       vertex 2.642168 5.000000 2.642167
       vertex 2.687290 5.000000 2.574636
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 -2.872500
       vertex -2.593326 5.000000 -2.593327
       vertex -2.744413 5.000000 -1.833761
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex -2.198516 5.000000 2.198516
       vertex -2.593327 5.000000 2.593327
       vertex -1.189828 5.000000 2.872500
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.593327 5.000000 2.593327
       vertex -2.198516 5.000000 2.198516
       vertex -2.744415 5.000000 1.833758
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.833760 5.000000 2.744414
       vertex 2.687290 5.000000 2.574636
       vertex 3.407036 5.000000 1.497460
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.833760 5.000000 2.744414
       vertex 3.407036 5.000000 1.497460
       vertex 3.237258 5.000000 0.643930
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.198516 5.000000 -2.198518
       vertex -2.744413 5.000000 -1.833761
       vertex -2.872499 5.000000 -1.189830
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189824 5.000000 2.872501
       vertex 1.833760 5.000000 2.744414
       vertex 2.198519 5.000000 2.198514
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.198517 5.000000 -2.198517
       vertex 2.872501 5.000000 -1.189826
       vertex 2.744415 5.000000 -1.833760
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 2.872500 5.000000 1.189829
       vertex 3.237258 5.000000 0.643930
       vertex 3.109172 5.000000 -0.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -2.198516 5.000000 2.198516
       vertex -2.872500 5.000000 1.189829
       vertex -2.744415 5.000000 1.833758
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -5.000000 4.103733
       vertex -5.000000 -5.000000 5.000000
       vertex -5.000000 -4.401134 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -4.597635 -5.000000
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 -5.000000 4.103733
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -5.000000 4.103733
       vertex -5.000000 -4.401134 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -4.401134 5.000000
       vertex -5.000000 -1.620707 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 5.000000 -0.027337
       vertex -5.000000 5.000000 -5.000000
       vertex -5.000000 0.027335 -5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -2.299216 -2.299216
       vertex -5.000000 -0.009603 -5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -0.009603 -5.000000
       vertex -5.000000 -4.597635 -5.000000
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 -2.299215 -2.299215
       vertex -5.000000 -4.427082 -4.427082
       vertex -5.000000 -3.405245 -0.994563
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.093314 -4.934021
       vertex -5.000000 0.027335 -5.000000
       vertex -5.000000 0.009602 -5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 4.882087 4.882087
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 5.000000 4.909753
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 4.882087 4.882087
       vertex -5.000000 -2.146763 3.232888
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 -2.146763 3.232888
       vertex -5.000000 -1.620707 5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.513671 2.513671
       vertex -5.000000 4.882087 4.882087
       vertex -5.000000 5.000000 4.909753
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.513671 2.513671
       vertex -5.000000 5.000000 4.909753
       vertex -5.000000 5.000000 0.027348
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 4.882087 4.882087
       vertex -5.000000 2.513672 2.513672
       vertex -5.000000 1.045454 3.981887
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.045450 -3.981886
       vertex -5.000000 0.093314 -4.934021
       vertex -5.000000 0.009602 -5.000000
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.045450 -3.981886
       vertex -5.000000 0.009602 -5.000000
       vertex -5.000000 -0.009603 -5.000000
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 1.045450 -3.981886
       vertex -5.000000 -0.009603 -5.000000
       vertex -5.000000 -1.349002 -3.420069
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.296098 2.731248
       vertex -5.000000 -2.146763 3.232888
       vertex -5.000000 -1.686450 3.340893
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.438503 2.252885
       vertex -5.000000 -2.296098 2.731248
       vertex -5.000000 -1.686450 3.340893
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.438503 2.252885
       vertex -5.000000 -1.686450 3.340893
       vertex -5.000000 1.045454 3.981887
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.438503 2.252885
       vertex -5.000000 1.045454 3.981887
       vertex -5.000000 1.419654 3.607688
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 2.071069
       vertex -5.000000 -2.438503 2.252885
       vertex -5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.299215 2.299215
       vertex -5.000000 2.513671 2.513671
       vertex -5.000000 5.000000 0.027348
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.299215 2.299215
       vertex -5.000000 5.000000 0.027348
       vertex -5.000000 5.000000 -0.027337
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.299215 2.299215
       vertex -5.000000 5.000000 -0.027337
       vertex -5.000000 4.605777 -0.421560
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 2.513672 2.513672
       vertex -5.000000 2.299215 2.299215
       vertex -5.000000 1.242640 3.545529
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.513672 2.513672
       vertex -5.000000 1.242640 3.545529
       vertex -5.000000 1.419654 3.607688
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 3.805954 0.521895
       vertex -5.000000 4.605777 -0.421560
       vertex -5.000000 3.538742 -1.488595
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.641431 0.715963
       vertex -5.000000 3.805954 0.521895
       vertex -5.000000 3.736590 0.000002
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.956266 -2.071070
       vertex -5.000000 1.045450 -3.981886
       vertex -5.000000 -1.071069 -3.485281
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071067
       vertex -5.000000 3.641431 0.715963
       vertex -5.000000 3.736590 0.000002
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071067
       vertex -5.000000 3.736590 0.000002
       vertex -5.000000 3.542932 -1.457072
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 1.071069 3.485281
       vertex -5.000000 1.242640 3.545529
       vertex -5.000000 1.349000 3.420069
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.000000 -3.109172
       vertex -5.000000 -1.071069 -3.485281
       vertex -5.000000 -1.349002 -3.420069
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 0.000000 -3.109172
       vertex -5.000000 -1.349002 -3.420069
       vertex -5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.000001 3.109172
       vertex -5.000000 1.071069 3.485281
       vertex -5.000000 1.349000 3.420069
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 0.000001 3.109172
       vertex -5.000000 1.349000 3.420069
       vertex -5.000000 2.296097 2.302892
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.109171 -0.000000
       vertex -5.000000 3.542932 -1.457072
       vertex -5.000000 3.538742 -1.488595
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 3.109171 -0.000000
       vertex -5.000000 3.538742 -1.488595
       vertex -5.000000 2.956266 -2.071070
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 3.109171 -0.000000
       vertex -5.000000 2.956266 -2.071070
       vertex -5.000000 2.438504 -2.252884
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071067
       vertex -5.000000 -3.405245 -0.994563
       vertex -5.000000 -3.109171 -0.000001
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 -2.071070
       vertex -5.000000 2.438504 -2.252884
       vertex -5.000000 2.296103 -2.302889
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.563426 -5.000000 -5.000000
       vertex 5.000000 -5.000000 -5.000000
       vertex 5.000000 -5.000000 -1.353408
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.593328 -5.000000 2.593328
       vertex 5.000000 -5.000000 5.000000
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.593328 -5.000000 2.593328
       vertex -5.000000 -5.000000 5.000000
       vertex -5.000000 -5.000000 4.103733
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 2.593328 -5.000000 2.593328
       vertex 5.000000 -5.000000 2.114611
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -4.664392 -5.000000 -4.664392
       vertex -2.999461 -5.000000 3.705801
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -2.999461 -5.000000 3.705801
       vertex -5.000000 -5.000000 4.103733
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -4.664392 -5.000000 -4.664392
       vertex -5.000000 -5.000000 -5.000000
       vertex -4.731149 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -4.664392 -5.000000 -4.664392
       vertex -2.642167 -5.000000 -2.642167
       vertex -3.890522 -5.000000 -0.773874
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.642167 -5.000000 -2.642167
       vertex -4.664392 -5.000000 -4.664392
       vertex -4.731149 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -2.642167 -5.000000 -2.642167
       vertex -4.731149 -5.000000 -5.000000
       vertex -1.066714 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.686719 -5.000000 -4.072097
       vertex -1.066714 -5.000000 -5.000000
       vertex 2.563426 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.686719 -5.000000 -4.072097
       vertex 2.563426 -5.000000 -5.000000
       vertex 3.929933 -5.000000 -2.954878
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 3.116650 -5.000000 -3.116650
       vertex 3.929933 -5.000000 -2.954878
       vertex 5.000000 -5.000000 -1.353408
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 3.116650 -5.000000 -3.116650
       vertex 5.000000 -5.000000 -1.353408
       vertex 5.000000 -5.000000 2.114611
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 3.116650 -5.000000 -3.116650
       vertex 5.000000 -5.000000 2.114611
       vertex 4.189287 -5.000000 2.275872
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex -1.497461 -5.000000 3.407035
       vertex -2.999461 -5.000000 3.705801
       vertex -3.298226 -5.000000 2.203802
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex -1.403498 -5.000000 3.388345
       vertex -1.497461 -5.000000 3.407035
       vertex -3.298226 -5.000000 2.203802
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.403498 -5.000000 3.388345
       vertex -3.298226 -5.000000 2.203802
       vertex -3.765867 -5.000000 -0.147192
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.203802 -5.000000 -3.298227
       vertex -1.686719 -5.000000 -4.072097
       vertex -0.000003 -5.000000 -3.736589
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.642167 -5.000000 -2.642167
       vertex -2.593328 -5.000000 -2.593328
       vertex -2.687289 -5.000000 -2.574637
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -2.593327 -5.000000 -2.593327
       vertex -2.642167 -5.000000 -2.642167
       vertex -2.203802 -5.000000 -3.298227
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.593327 -5.000000 -2.593327
       vertex -2.203802 -5.000000 -3.298227
       vertex -0.000003 -5.000000 -3.736589
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.593327 -5.000000 -2.593327
       vertex -0.000003 -5.000000 -3.736589
       vertex 1.577118 -5.000000 -3.422881
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.593328 -5.000000 -2.593328
       vertex -2.198517 -5.000000 -2.198517
       vertex -3.535534 -5.000000 -1.305150
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.593328 -5.000000 -2.593328
       vertex -3.535534 -5.000000 -1.305150
       vertex -2.687289 -5.000000 -2.574637
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -2.198517 -5.000000 -2.198517
       vertex -2.593327 -5.000000 -2.593327
       vertex -1.189827 -5.000000 -2.872501
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.198516 -5.000000 2.198516
       vertex 2.593328 -5.000000 2.593328
       vertex 1.833759 -5.000000 2.744415
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex 2.593328 -5.000000 2.593328
       vertex 2.198516 -5.000000 2.198516
       vertex 3.720743 -5.000000 -0.079659
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.593328 -5.000000 2.593328
       vertex 3.720743 -5.000000 -0.079659
       vertex 4.189287 -5.000000 2.275872
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 3.422880 -5.000000 -1.577121
       vertex 3.720743 -5.000000 -0.079659
       vertex 2.872501 -5.000000 1.189826
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.872501 -5.000000 1.189826
       vertex -3.765867 -5.000000 -0.147192
       vertex -3.890522 -5.000000 -0.773874
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.872501 -5.000000 1.189826
       vertex -3.890522 -5.000000 -0.773874
       vertex -3.535534 -5.000000 -1.305150
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.872501 -5.000000 1.189826
       vertex -3.535534 -5.000000 -1.305150
       vertex -3.388345 -5.000000 -1.403498
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.643930 -5.000000 -3.237258
       vertex 1.577118 -5.000000 -3.422881
       vertex 3.116650 -5.000000 -3.116650
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.643930 -5.000000 -3.237258
       vertex 3.116650 -5.000000 -3.116650
       vertex 3.422880 -5.000000 -1.577121
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 0.643930 -5.000000 -3.237258
       vertex 3.422880 -5.000000 -1.577121
       vertex 3.388345 -5.000000 -1.403499
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex 0.000001 -5.000000 3.109172
       vertex -1.403498 -5.000000 3.388345
       vertex -1.833758 -5.000000 2.744415
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.744416 -5.000000 -1.833758
       vertex 3.388345 -5.000000 -1.403499
       vertex 3.109173 -5.000000 -0.000003
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.198518 -5.000000 -2.198515
       vertex 2.744416 -5.000000 -1.833758
       vertex 2.872503 -5.000000 -1.189823
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189828 -5.000000 2.872500
       vertex -1.833758 -5.000000 2.744415
       vertex -2.198515 -5.000000 2.198518
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -2.198517 -5.000000 -2.198517
       vertex -3.237257 -5.000000 -0.643930
       vertex -3.388345 -5.000000 -1.403498
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -0.000001 -5.000000 -3.109172
       vertex 0.643930 -5.000000 -3.237258
       vertex 1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex -3.109172 -5.000000 0.000001
       vertex -3.237257 -5.000000 -0.643930
       vertex -2.872501 -5.000000 -1.189828
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 2.198516 -5.000000 2.198516
       vertex 1.833759 -5.000000 2.744415
       vertex 1.189827 -5.000000 2.872501
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.882087 4.882087 -5.000000
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 -1.000000
    outer loop
       vertex 4.882087 4.882087 -5.000000
       vertex 5.000000 -5.000000 -5.000000
       vertex 2.563426 -5.000000 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 4.882087 4.882087 -5.000000
       vertex 4.909754 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -4.838835 -4.838835 -5.000000
       vertex -4.731149 -5.000000 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -4.838835 -4.838835 -5.000000
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 -4.597635 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.652134 4.652134 -5.000000
       vertex 4.882087 4.882087 -5.000000
       vertex 4.821434 4.623587 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 4.882087 4.882087 -5.000000
       vertex 4.652135 4.652135 -5.000000
       vertex 2.589250 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.882087 4.882087 -5.000000
       vertex 2.589250 5.000000 -5.000000
       vertex 4.909754 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -5.000000 0.027337 -5.000000
       vertex -5.000000 5.000000 -5.000000
       vertex -4.010875 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.079602 3.079602 -5.000000
       vertex 4.652134 4.652134 -5.000000
       vertex 4.821434 4.623587 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.079602 3.079602 -5.000000
       vertex 4.821434 4.623587 -5.000000
       vertex 4.346643 2.600037 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 4.652135 4.652135 -5.000000
       vertex 3.079603 3.079603 -5.000000
       vertex -1.994201 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 4.652135 4.652135 -5.000000
       vertex -1.994201 5.000000 -5.000000
       vertex 2.589250 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -4.838835 -4.838835 -5.000000
       vertex -2.800386 -2.800386 -5.000000
       vertex -1.066715 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -4.838835 -4.838835 -5.000000
       vertex -1.066715 -5.000000 -5.000000
       vertex -4.731149 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.800386 -2.800386 -5.000000
       vertex -4.838835 -4.838835 -5.000000
       vertex -5.000000 -4.597635 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -2.800386 -2.800386 -5.000000
       vertex -5.000000 -4.597635 -5.000000
       vertex -5.000000 -0.009604 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.232888 -2.146763 -5.000000
       vertex 2.563426 -5.000000 -5.000000
       vertex -1.066715 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.232888 -2.146763 -5.000000
       vertex -1.066715 -5.000000 -5.000000
       vertex -2.071070 -3.725714 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -5.000000 0.009603 -5.000000
       vertex -5.000000 0.027337 -5.000000
       vertex -4.010875 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -5.000000 0.009603 -5.000000
       vertex -4.010875 5.000000 -5.000000
       vertex -1.994201 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -5.000000 0.009603 -5.000000
       vertex -1.994201 5.000000 -5.000000
       vertex -0.002599 4.246193 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.800386 -2.800386 -5.000000
       vertex -2.301126 -2.301126 -5.000000
       vertex -1.826449 -3.652892 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.800386 -2.800386 -5.000000
       vertex -1.826449 -3.652892 -5.000000
       vertex -2.071070 -3.725714 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -2.301126 -2.301126 -5.000000
       vertex -2.800386 -2.800386 -5.000000
       vertex -5.000000 -0.009604 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.301126 -2.301126 -5.000000
       vertex -5.000000 -0.009604 -5.000000
       vertex -5.000000 0.009603 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.301126 -2.301126 -5.000000
       vertex -5.000000 0.009603 -5.000000
       vertex -3.545530 1.242643 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -0.521896 3.805954 -5.000000
       vertex -0.002599 4.246193 -5.000000
       vertex 2.071065 3.461327 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.301125 2.301125 -5.000000
       vertex 3.079602 3.079602 -5.000000
       vertex 4.346643 2.600037 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.301125 2.301125 -5.000000
       vertex 4.346643 2.600037 -5.000000
       vertex 3.485281 -1.071072 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 3.079603 3.079603 -5.000000
       vertex 2.301124 2.301124 -5.000000
       vertex 1.885034 3.486052 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.079603 3.079603 -5.000000
       vertex 1.885034 3.486052 -5.000000
       vertex 2.071065 3.461327 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -0.000001 -3.109171 -5.000000
       vertex -1.826449 -3.652892 -5.000000
       vertex -2.252885 -2.438504 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.301126 -2.301126 -5.000000
       vertex -2.299216 -2.299216 -5.000000
       vertex -2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.301126 -2.301126 -5.000000
       vertex -2.071068 -2.492630 -5.000000
       vertex -2.252885 -2.438504 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex -2.299215 -2.299215 -5.000000
       vertex -2.301126 -2.301126 -5.000000
       vertex -2.302890 -2.296100 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -0.715964 3.641431 -5.000000
       vertex -0.521896 3.805954 -5.000000
       vertex -0.000000 3.736590 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.485281 1.071070 -5.000000
       vertex -3.545530 1.242643 -5.000000
       vertex -3.420069 1.349003 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.340894 -1.686447 -5.000000
       vertex 3.232888 -2.146763 -5.000000
       vertex 2.731242 -2.296100 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.109171 -0.000002 -5.000000
       vertex 3.485281 -1.071072 -5.000000
       vertex 3.340894 -1.686447 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 -1.000000
    outer loop
       vertex 3.109171 -0.000002 -5.000000
       vertex 3.340894 -1.686447 -5.000000
       vertex 2.731242 -2.296100 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 3.109171 -0.000002 -5.000000
       vertex 2.731242 -2.296100 -5.000000
       vertex 2.252885 -2.438503 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.071071 2.492628 -5.000000
       vertex -0.715964 3.641431 -5.000000
       vertex -0.000000 3.736590 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.071071 2.492628 -5.000000
       vertex -0.000000 3.736590 -5.000000
       vertex 1.457075 3.542931 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 1.488595 3.538742 -5.000000
       vertex 1.885034 3.486052 -5.000000
       vertex 2.071069 2.956268 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 0.000001 3.109171 -5.000000
       vertex 1.457075 3.542931 -5.000000
       vertex 1.488595 3.538742 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 0.000001 3.109171 -5.000000
       vertex 1.488595 3.538742 -5.000000
       vertex 2.071069 2.956268 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 0.000001 3.109171 -5.000000
       vertex 2.071069 2.956268 -5.000000
       vertex 2.252883 2.438504 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.109172 0.000000 -5.000000
       vertex -3.485281 1.071070 -5.000000
       vertex -3.420069 1.349003 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -3.109172 0.000000 -5.000000
       vertex -3.420069 1.349003 -5.000000
       vertex -2.302890 2.296102 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.299216 2.299216 -5.000000
       vertex 2.301125 2.301125 -5.000000
       vertex 2.302888 2.296103 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 2.301124 2.301124 -5.000000
       vertex 2.299216 2.299216 -5.000000
       vertex 2.071071 2.492629 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.301124 2.301124 -5.000000
       vertex 2.071071 2.492629 -5.000000
       vertex 2.252883 2.438504 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302891 -2.296099 -5.000000
       vertex 2.252885 -2.438503 -5.000000
       vertex 2.071067 -2.492629 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695438
    outer loop
       vertex -2.071068 2.492630 -5.000000
       vertex -2.121320 2.296101 -5.121321
       vertex -2.171573 2.407425 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex -2.171573 2.407425 -5.000000
       vertex -2.121320 2.296101 -5.121321
       vertex -2.302890 2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex 2.121320 -2.296101 -5.121321
       vertex 2.302890 -2.296101 -5.000000
       vertex 2.171573 -2.407425 -5.000000
    endloop
  endfacet
  facet normal 0.464678 -0.548124 -0.695438
    outer loop
       vertex 2.121320 -2.296101 -5.121321
       vertex 2.171573 -2.407425 -5.000000
       vertex 2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.191481 -0.191481 -0.962637
    outer loop
       vertex -2.296101 0.000000 -5.543277
       vertex -0.000000 0.000000 -6.000000
       vertex -0.000000 -2.296101 -5.543277
    endloop
  endfacet
  facet normal -0.191480 -0.191480 -0.962637
    outer loop
       vertex -2.296101 0.000000 -5.543277
       vertex -0.000000 -2.296101 -5.543277
       vertex -2.121320 -2.296101 -5.121321
    endloop
  endfacet
  facet normal -0.464677 -0.548124 -0.695439
    outer loop
       vertex -2.302890 -2.296101 -5.000000
       vertex -2.121320 -2.296101 -5.121321
       vertex -2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.191480 0.191481 -0.962637
    outer loop
       vertex -2.121320 2.296101 -5.121321
       vertex -0.000000 2.296101 -5.543277
       vertex -0.000000 0.000000 -6.000000
    endloop
  endfacet
  facet normal -0.191481 0.191480 -0.962637
    outer loop
       vertex -2.121320 2.296101 -5.121321
       vertex -0.000000 0.000000 -6.000000
       vertex -2.296101 0.000000 -5.543277
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816084
    outer loop
       vertex 2.296101 0.000000 -5.543277
       vertex 3.109172 0.000000 -5.000000
       vertex 2.839378 -0.768310 -5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex 2.296101 0.000000 -5.543277
       vertex 2.839378 -0.768310 -5.000000
       vertex 2.302890 -2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex 2.296101 0.000000 -5.543277
       vertex 2.302890 -2.296101 -5.000000
       vertex 2.121320 -2.296101 -5.121321
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex 2.121320 2.296101 -5.121321
       vertex 2.302890 2.296101 -5.000000
       vertex 2.414214 1.979074 -5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816084
    outer loop
       vertex 2.121320 2.296101 -5.121321
       vertex 2.414214 1.979074 -5.000000
       vertex 3.109172 0.000000 -5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex 2.121320 2.296101 -5.121321
       vertex 3.109172 0.000000 -5.000000
       vertex 2.296101 0.000000 -5.543277
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex -0.000000 3.109172 -5.000000
       vertex -0.000000 2.296101 -5.543277
       vertex -0.543277 2.947442 -5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex -0.543277 2.947442 -5.000000
       vertex -0.000000 2.296101 -5.543277
       vertex -2.121320 2.296101 -5.121321
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex -0.543277 2.947442 -5.000000
       vertex -2.121320 2.296101 -5.121321
       vertex -2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex -2.121320 -2.296101 -5.121321
       vertex -0.000000 -2.296101 -5.543277
       vertex -0.000000 -3.109172 -5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex -2.121320 -2.296101 -5.121321
       vertex -0.000000 -3.109172 -5.000000
       vertex -1.828426 -2.564862 -5.000000
    endloop
  endfacet
  facet normal -0.163172 -0.548124 -0.820326
    outer loop
       vertex -2.121320 -2.296101 -5.121321
       vertex -1.828426 -2.564862 -5.000000
       vertex -2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.191480 0.191480 -0.962637
    outer loop
       vertex -0.000000 2.296101 -5.543277
       vertex 2.121320 2.296101 -5.121321
       vertex 2.296101 0.000000 -5.543277
    endloop
  endfacet
  facet normal 0.191481 0.191481 -0.962637
    outer loop
       vertex -0.000000 2.296101 -5.543277
       vertex 2.296101 0.000000 -5.543277
       vertex -0.000000 0.000000 -6.000000
    endloop
  endfacet
  facet normal 0.163172 0.548124 -0.820326
    outer loop
       vertex 2.071068 2.492630 -5.000000
       vertex 2.121320 2.296101 -5.121321
       vertex 1.828426 2.564862 -5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex 1.828426 2.564862 -5.000000
       vertex 2.121320 2.296101 -5.121321
       vertex -0.000000 2.296101 -5.543277
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex 1.828426 2.564862 -5.000000
       vertex -0.000000 2.296101 -5.543277
       vertex -0.000000 3.109172 -5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex 2.302890 2.296101 -5.000000
       vertex 2.121320 2.296101 -5.121321
       vertex 2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816084
    outer loop
       vertex -3.109172 0.000000 -5.000000
       vertex -2.296101 0.000000 -5.543277
       vertex -2.121320 -2.296101 -5.121321
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816084
    outer loop
       vertex -3.109172 0.000000 -5.000000
       vertex -2.121320 -2.296101 -5.121321
       vertex -2.414214 -1.979074 -5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816085
    outer loop
       vertex -2.414214 -1.979074 -5.000000
       vertex -2.121320 -2.296101 -5.121321
       vertex -2.302890 -2.296101 -5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex -2.302890 2.296101 -5.000000
       vertex -2.121320 2.296101 -5.121321
       vertex -2.296101 0.000000 -5.543277
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex -2.302890 2.296101 -5.000000
       vertex -2.296101 0.000000 -5.543277
       vertex -2.839378 0.768310 -5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816084
    outer loop
       vertex -2.839378 0.768310 -5.000000
       vertex -2.296101 0.000000 -5.543277
       vertex -3.109172 0.000000 -5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex -0.000000 -2.296101 -5.543277
       vertex 2.121320 -2.296101 -5.121321
       vertex 2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex -0.000000 -2.296101 -5.543277
       vertex 2.071068 -2.492630 -5.000000
       vertex 0.543277 -2.947442 -5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex -0.000000 -2.296101 -5.543277
       vertex 0.543277 -2.947442 -5.000000
       vertex -0.000000 -3.109172 -5.000000
    endloop
  endfacet
  facet normal 0.191481 -0.191480 -0.962637
    outer loop
       vertex -0.000000 0.000000 -6.000000
       vertex 2.296101 0.000000 -5.543277
       vertex 2.121320 -2.296101 -5.121321
    endloop
  endfacet
  facet normal 0.191480 -0.191481 -0.962637
    outer loop
       vertex -0.000000 0.000000 -6.000000
       vertex 2.121320 -2.296101 -5.121321
       vertex -0.000000 -2.296101 -5.543277
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552209
    outer loop
       vertex -0.000000 -5.000000 -3.109172
       vertex -0.000000 -5.543277 -2.296101
       vertex -0.678174 -5.000000 -2.974275
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552208
    outer loop
       vertex -0.678174 -5.000000 -2.974275
       vertex -0.000000 -5.543277 -2.296101
       vertex -0.878680 -5.543277 -2.121320
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552208
    outer loop
       vertex -0.678174 -5.000000 -2.974275
       vertex -0.878680 -5.543277 -2.121320
       vertex -1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex -1.189829 -5.000000 -2.872500
       vertex -0.878680 -5.543277 -2.121320
       vertex -1.764757 -5.000000 -2.488345
    endloop
  endfacet
  facet normal -0.312801 -0.826438 -0.468140
    outer loop
       vertex -1.764757 -5.000000 -2.488345
       vertex -0.878680 -5.543277 -2.121320
       vertex -1.623588 -5.543277 -1.623588
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex -1.764757 -5.000000 -2.488345
       vertex -1.623588 -5.543277 -1.623588
       vertex -2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.194944 -0.980048 0.038777
    outer loop
       vertex 2.296101 -5.543277 0.000000
       vertex 2.121320 -5.543277 0.878680
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex 2.198517 -5.000000 -2.198517
       vertex 1.623588 -5.543277 -1.623588
       vertex 1.623588 -5.000000 -2.582671
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex 1.623588 -5.000000 -2.582671
       vertex 1.623588 -5.543277 -1.623588
       vertex 0.878680 -5.543277 -2.121320
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex 1.623588 -5.000000 -2.582671
       vertex 0.878680 -5.543277 -2.121320
       vertex 1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.194944 -0.980048 -0.038777
    outer loop
       vertex 2.121320 -5.543277 -0.878680
       vertex 2.296101 -5.543277 -0.000000
       vertex 0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552208
    outer loop
       vertex 1.189829 -5.000000 -2.872500
       vertex 0.878680 -5.543277 -2.121320
       vertex 0.511654 -5.000000 -3.007397
    endloop
  endfacet
  facet normal 0.109841 -0.826438 -0.552209
    outer loop
       vertex 0.511654 -5.000000 -3.007397
       vertex 0.878680 -5.543277 -2.121320
       vertex -0.000000 -5.543277 -2.296101
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552209
    outer loop
       vertex 0.511654 -5.000000 -3.007397
       vertex -0.000000 -5.543277 -2.296101
       vertex -0.000000 -5.000000 -3.109172
    endloop
  endfacet
  facet normal -0.552208 -0.826439 -0.109841
    outer loop
       vertex -2.872500 -5.000000 -1.189829
       vertex -2.121320 -5.543277 -0.878680
       vertex -3.007397 -5.000000 -0.511654
    endloop
  endfacet
  facet normal -0.552209 -0.826438 -0.109841
    outer loop
       vertex -3.007397 -5.000000 -0.511654
       vertex -2.121320 -5.543277 -0.878680
       vertex -2.296101 -5.543277 0.000000
    endloop
  endfacet
  facet normal -0.552209 -0.826439 -0.109841
    outer loop
       vertex -3.007397 -5.000000 -0.511654
       vertex -2.296101 -5.543277 0.000000
       vertex -3.109172 -5.000000 0.000000
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex 1.189829 -5.000000 2.872500
       vertex 0.878680 -5.543277 2.121320
       vertex 1.764757 -5.000000 2.488345
    endloop
  endfacet
  facet normal 0.312801 -0.826438 0.468140
    outer loop
       vertex 1.764757 -5.000000 2.488345
       vertex 0.878680 -5.543277 2.121320
       vertex 1.623588 -5.543277 1.623588
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex 1.764757 -5.000000 2.488345
       vertex 1.623588 -5.543277 1.623588
       vertex 2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex 2.198517 -5.000000 2.198517
       vertex 1.623588 -5.543277 1.623588
       vertex 2.582671 -5.000000 1.623588
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex 2.582671 -5.000000 1.623588
       vertex 1.623588 -5.543277 1.623588
       vertex 2.121320 -5.543277 0.878680
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex 2.582671 -5.000000 1.623588
       vertex 2.121320 -5.543277 0.878680
       vertex 2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal 0.165265 -0.980048 0.110427
    outer loop
       vertex 2.121320 -5.543277 0.878680
       vertex 1.623588 -5.543277 1.623588
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex -2.198517 -5.000000 -2.198517
       vertex -1.623588 -5.543277 -1.623588
       vertex -2.582671 -5.000000 -1.623588
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex -2.582671 -5.000000 -1.623588
       vertex -1.623588 -5.543277 -1.623588
       vertex -2.121320 -5.543277 -0.878680
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex -2.582671 -5.000000 -1.623588
       vertex -2.121320 -5.543277 -0.878680
       vertex -2.872500 -5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552209
    outer loop
       vertex 0.000000 -5.000000 3.109172
       vertex 0.000000 -5.543277 2.296101
       vertex 0.678174 -5.000000 2.974275
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552208
    outer loop
       vertex 0.678174 -5.000000 2.974275
       vertex 0.000000 -5.543277 2.296101
       vertex 0.878680 -5.543277 2.121320
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552208
    outer loop
       vertex 0.678174 -5.000000 2.974275
       vertex 0.878680 -5.543277 2.121320
       vertex 1.189829 -5.000000 2.872500
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex -2.872500 -5.000000 1.189829
       vertex -2.121320 -5.543277 0.878680
       vertex -2.488345 -5.000000 1.764757
    endloop
  endfacet
  facet normal -0.468140 -0.826438 0.312801
    outer loop
       vertex -2.488345 -5.000000 1.764757
       vertex -2.121320 -5.543277 0.878680
       vertex -1.623588 -5.543277 1.623588
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex -2.488345 -5.000000 1.764757
       vertex -1.623588 -5.543277 1.623588
       vertex -2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552208
    outer loop
       vertex -1.189829 -5.000000 2.872500
       vertex -0.878680 -5.543277 2.121320
       vertex -0.511654 -5.000000 3.007397
    endloop
  endfacet
  facet normal -0.109841 -0.826438 0.552209
    outer loop
       vertex -0.511654 -5.000000 3.007397
       vertex -0.878680 -5.543277 2.121320
       vertex 0.000000 -5.543277 2.296101
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552209
    outer loop
       vertex -0.511654 -5.000000 3.007397
       vertex 0.000000 -5.543277 2.296101
       vertex 0.000000 -5.000000 3.109172
    endloop
  endfacet
  facet normal 0.038777 -0.980048 -0.194944
    outer loop
       vertex -0.000000 -5.543277 -2.296101
       vertex 0.878680 -5.543277 -2.121320
       vertex -0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal 0.165265 -0.980048 -0.110427
    outer loop
       vertex 1.623588 -5.543277 -1.623588
       vertex 2.121320 -5.543277 -0.878680
       vertex 0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal -0.552209 -0.826439 0.109841
    outer loop
       vertex -3.109172 -5.000000 0.000000
       vertex -2.296101 -5.543277 0.000000
       vertex -2.974275 -5.000000 0.678174
    endloop
  endfacet
  facet normal -0.552208 -0.826439 0.109841
    outer loop
       vertex -2.974275 -5.000000 0.678174
       vertex -2.296101 -5.543277 0.000000
       vertex -2.121320 -5.543277 0.878680
    endloop
  endfacet
  facet normal -0.552208 -0.826439 0.109841
    outer loop
       vertex -2.974275 -5.000000 0.678174
       vertex -2.121320 -5.543277 0.878680
       vertex -2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal -0.194944 -0.980048 -0.038777
    outer loop
       vertex -2.296101 -5.543277 0.000000
       vertex -2.121320 -5.543277 -0.878680
       vertex -0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex 2.872500 -5.000000 -1.189829
       vertex 2.121320 -5.543277 -0.878680
       vertex 2.488345 -5.000000 -1.764757
    endloop
  endfacet
  facet normal 0.468140 -0.826438 -0.312801
    outer loop
       vertex 2.488345 -5.000000 -1.764757
       vertex 2.121320 -5.543277 -0.878680
       vertex 1.623588 -5.543277 -1.623588
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex 2.488345 -5.000000 -1.764757
       vertex 1.623588 -5.543277 -1.623588
       vertex 2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.110427 -0.980048 -0.165265
    outer loop
       vertex 0.878680 -5.543277 -2.121320
       vertex 1.623588 -5.543277 -1.623588
       vertex 0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal -0.038777 -0.980048 -0.194944
    outer loop
       vertex -0.878680 -5.543277 -2.121320
       vertex -0.000000 -5.543277 -2.296101
       vertex -0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal -0.165265 -0.980048 0.110427
    outer loop
       vertex -1.623588 -5.543277 1.623588
       vertex -2.121320 -5.543277 0.878680
       vertex -0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.552208 -0.826439 0.109841
    outer loop
       vertex 2.872500 -5.000000 1.189829
       vertex 2.121320 -5.543277 0.878680
       vertex 3.007397 -5.000000 0.511654
    endloop
  endfacet
  facet normal 0.552209 -0.826438 0.109841
    outer loop
       vertex 3.007397 -5.000000 0.511654
       vertex 2.121320 -5.543277 0.878680
       vertex 2.296101 -5.543277 0.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 0.109841
    outer loop
       vertex 3.007397 -5.000000 0.511654
       vertex 2.296101 -5.543277 0.000000
       vertex 3.109172 -5.000000 0.000000
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex -2.198517 -5.000000 2.198517
       vertex -1.623588 -5.543277 1.623588
       vertex -1.623588 -5.000000 2.582671
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex -1.623588 -5.000000 2.582671
       vertex -1.623588 -5.543277 1.623588
       vertex -0.878680 -5.543277 2.121320
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex -1.623588 -5.000000 2.582671
       vertex -0.878680 -5.543277 2.121320
       vertex -1.189829 -5.000000 2.872500
    endloop
  endfacet
  facet normal -0.110427 -0.980048 0.165265
    outer loop
       vertex -0.878680 -5.543277 2.121320
       vertex -1.623588 -5.543277 1.623588
       vertex -0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.110427 -0.980048 0.165265
    outer loop
       vertex 1.623588 -5.543277 1.623588
       vertex 0.878680 -5.543277 2.121320
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 -0.109841
    outer loop
       vertex 3.109172 -5.000000 -0.000000
       vertex 2.296101 -5.543277 -0.000000
       vertex 2.974275 -5.000000 -0.678174
    endloop
  endfacet
  facet normal 0.552208 -0.826439 -0.109841
    outer loop
       vertex 2.974275 -5.000000 -0.678174
       vertex 2.296101 -5.543277 -0.000000
       vertex 2.121320 -5.543277 -0.878680
    endloop
  endfacet
  facet normal 0.552208 -0.826439 -0.109841
    outer loop
       vertex 2.974275 -5.000000 -0.678174
       vertex 2.121320 -5.543277 -0.878680
       vertex 2.872500 -5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.194944 -0.980048 0.038777
    outer loop
       vertex -2.121320 -5.543277 0.878680
       vertex -2.296101 -5.543277 0.000000
       vertex -0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.038777 -0.980048 0.194944
    outer loop
       vertex 0.878680 -5.543277 2.121320
       vertex 0.000000 -5.543277 2.296101
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal -0.165265 -0.980048 -0.110427
    outer loop
       vertex -2.121320 -5.543277 -0.878680
       vertex -1.623588 -5.543277 -1.623588
       vertex -0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal -0.110427 -0.980048 -0.165265
    outer loop
       vertex -1.623588 -5.543277 -1.623588
       vertex -0.878680 -5.543277 -2.121320
       vertex -0.000000 -6.000000 -0.000000
    endloop
  endfacet
  facet normal -0.038777 -0.980048 0.194944
    outer loop
       vertex 0.000000 -5.543277 2.296101
       vertex -0.878680 -5.543277 2.121320
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.543277 -2.296101 0.000000
       vertex -5.121321 -2.296101 -2.121320
       vertex -5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.543277 -2.296101 0.000000
       vertex -5.000000 -2.492630 -2.071068
       vertex -5.000000 -2.947442 -0.543277
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.543277 -2.296101 0.000000
       vertex -5.000000 -2.947442 -0.543277
       vertex -5.000000 -3.109172 0.000000
    endloop
  endfacet
  facet normal -0.816084 -0.191480 0.545290
    outer loop
       vertex -5.000000 0.000000 3.109172
       vertex -5.543277 0.000000 2.296101
       vertex -5.121321 -2.296101 2.121320
    endloop
  endfacet
  facet normal -0.816084 -0.191480 0.545290
    outer loop
       vertex -5.000000 0.000000 3.109172
       vertex -5.121321 -2.296101 2.121320
       vertex -5.000000 -1.979074 2.414214
    endloop
  endfacet
  facet normal -0.816085 -0.191480 0.545290
    outer loop
       vertex -5.000000 -1.979074 2.414214
       vertex -5.121321 -2.296101 2.121320
       vertex -5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex -5.000000 2.296101 -2.302890
       vertex -5.121321 2.296101 -2.121320
       vertex -5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex -5.121321 -2.296101 -2.121320
       vertex -5.000000 -2.296101 -2.302890
       vertex -5.000000 -2.407425 -2.171573
    endloop
  endfacet
  facet normal -0.695438 -0.548124 -0.464678
    outer loop
       vertex -5.121321 -2.296101 -2.121320
       vertex -5.000000 -2.407425 -2.171573
       vertex -5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163172
    outer loop
       vertex -5.000000 2.492630 -2.071068
       vertex -5.121321 2.296101 -2.121320
       vertex -5.000000 2.564862 -1.828426
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex -5.000000 2.564862 -1.828426
       vertex -5.121321 2.296101 -2.121320
       vertex -5.543277 2.296101 0.000000
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex -5.000000 2.564862 -1.828426
       vertex -5.543277 2.296101 0.000000
       vertex -5.000000 3.109172 0.000000
    endloop
  endfacet
  facet normal -0.816084 -0.191480 -0.545290
    outer loop
       vertex -5.543277 0.000000 -2.296101
       vertex -5.000000 0.000000 -3.109172
       vertex -5.000000 -0.768310 -2.839378
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex -5.543277 0.000000 -2.296101
       vertex -5.000000 -0.768310 -2.839378
       vertex -5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex -5.543277 0.000000 -2.296101
       vertex -5.000000 -2.296101 -2.302890
       vertex -5.121321 -2.296101 -2.121320
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex -5.121321 -2.296101 2.121320
       vertex -5.543277 -2.296101 0.000000
       vertex -5.000000 -3.109172 0.000000
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex -5.121321 -2.296101 2.121320
       vertex -5.000000 -3.109172 0.000000
       vertex -5.000000 -2.564862 1.828426
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163172
    outer loop
       vertex -5.121321 -2.296101 2.121320
       vertex -5.000000 -2.564862 1.828426
       vertex -5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex -5.000000 2.296101 2.302890
       vertex -5.121321 2.296101 2.121320
       vertex -5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex -5.000000 2.296101 2.302890
       vertex -5.543277 0.000000 2.296101
       vertex -5.000000 0.768310 2.839378
    endloop
  endfacet
  facet normal -0.816084 0.191480 0.545290
    outer loop
       vertex -5.000000 0.768310 2.839378
       vertex -5.543277 0.000000 2.296101
       vertex -5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal -0.962637 0.191480 -0.191480
    outer loop
       vertex -5.543277 2.296101 0.000000
       vertex -5.121321 2.296101 -2.121320
       vertex -5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal -0.962637 0.191481 -0.191481
    outer loop
       vertex -5.543277 2.296101 0.000000
       vertex -5.543277 0.000000 -2.296101
       vertex -6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal -0.962637 -0.191480 -0.191481
    outer loop
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 0.000000 -2.296101
       vertex -5.121321 -2.296101 -2.121320
    endloop
  endfacet
  facet normal -0.962637 -0.191481 -0.191480
    outer loop
       vertex -6.000000 0.000000 0.000000
       vertex -5.121321 -2.296101 -2.121320
       vertex -5.543277 -2.296101 0.000000
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464677
    outer loop
       vertex -5.000000 -2.296101 2.302890
       vertex -5.121321 -2.296101 2.121320
       vertex -5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex -5.121321 2.296101 -2.121320
       vertex -5.000000 2.296101 -2.302890
       vertex -5.000000 1.979074 -2.414214
    endloop
  endfacet
  facet normal -0.816084 0.191480 -0.545290
    outer loop
       vertex -5.121321 2.296101 -2.121320
       vertex -5.000000 1.979074 -2.414214
       vertex -5.000000 0.000000 -3.109172
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex -5.121321 2.296101 -2.121320
       vertex -5.000000 0.000000 -3.109172
       vertex -5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex -5.000000 3.109172 0.000000
       vertex -5.543277 2.296101 0.000000
       vertex -5.000000 2.947442 0.543277
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex -5.000000 2.947442 0.543277
       vertex -5.543277 2.296101 0.000000
       vertex -5.121321 2.296101 2.121320
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex -5.000000 2.947442 0.543277
       vertex -5.121321 2.296101 2.121320
       vertex -5.000000 2.492630 2.071068
    endloop
  endfacet
  facet normal -0.962637 0.191481 0.191480
    outer loop
       vertex -5.121321 2.296101 2.121320
       vertex -5.543277 2.296101 0.000000
       vertex -6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal -0.962637 0.191480 0.191481
    outer loop
       vertex -5.121321 2.296101 2.121320
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal -0.962637 -0.191481 0.191481
    outer loop
       vertex -5.543277 0.000000 2.296101
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 -2.296101 0.000000
    endloop
  endfacet
  facet normal -0.962637 -0.191480 0.191480
    outer loop
       vertex -5.543277 0.000000 2.296101
       vertex -5.543277 -2.296101 0.000000
       vertex -5.121321 -2.296101 2.121320
    endloop
  endfacet
  facet normal -0.695438 0.548124 0.464677
    outer loop
       vertex -5.000000 2.492630 2.071068
       vertex -5.121321 2.296101 2.121320
       vertex -5.000000 2.407425 2.171573
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex -5.000000 2.407425 2.171573
       vertex -5.121321 2.296101 2.121320
       vertex -5.000000 2.296101 2.302890
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 5.543277 1.623588
       vertex 0.878680 5.543277 2.121320
       vertex 1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 5.543277 1.623588
       vertex 1.189829 5.000000 2.872500
       vertex 1.623588 5.000000 2.582671
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 5.543277 1.623588
       vertex 1.623588 5.000000 2.582671
       vertex 2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal 0.552209 0.826439 -0.109841
    outer loop
       vertex 2.121320 5.543277 -0.878680
       vertex 2.296101 5.543277 -0.000000
       vertex 3.109172 5.000000 -0.000000
    endloop
  endfacet
  facet normal 0.552209 0.826438 -0.109841
    outer loop
       vertex 2.121320 5.543277 -0.878680
       vertex 3.109172 5.000000 -0.000000
       vertex 3.007397 5.000000 -0.511654
    endloop
  endfacet
  facet normal 0.552208 0.826439 -0.109841
    outer loop
       vertex 2.121320 5.543277 -0.878680
       vertex 3.007397 5.000000 -0.511654
       vertex 2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.165265 0.980048 0.110427
    outer loop
       vertex -0.000000 6.000000 0.000000
       vertex -2.121320 5.543277 0.878680
       vertex -1.623588 5.543277 1.623588
    endloop
  endfacet
  facet normal 0.468140 0.826439 0.312801
    outer loop
       vertex 2.121320 5.543277 0.878680
       vertex 1.623588 5.543277 1.623588
       vertex 2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal 0.468140 0.826438 0.312801
    outer loop
       vertex 2.121320 5.543277 0.878680
       vertex 2.198517 5.000000 2.198517
       vertex 2.488345 5.000000 1.764757
    endloop
  endfacet
  facet normal 0.468140 0.826439 0.312801
    outer loop
       vertex 2.121320 5.543277 0.878680
       vertex 2.488345 5.000000 1.764757
       vertex 2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal -0.552209 0.826439 0.109841
    outer loop
       vertex -2.121320 5.543277 0.878680
       vertex -2.296101 5.543277 0.000000
       vertex -3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal -0.552209 0.826438 0.109841
    outer loop
       vertex -2.121320 5.543277 0.878680
       vertex -3.109172 5.000000 0.000000
       vertex -3.007397 5.000000 0.511654
    endloop
  endfacet
  facet normal -0.552208 0.826439 0.109841
    outer loop
       vertex -2.121320 5.543277 0.878680
       vertex -3.007397 5.000000 0.511654
       vertex -2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 5.543277 -1.623588
       vertex -0.878680 5.543277 -2.121320
       vertex -1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 5.543277 -1.623588
       vertex -1.189829 5.000000 -2.872500
       vertex -1.623588 5.000000 -2.582671
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 5.543277 -1.623588
       vertex -1.623588 5.000000 -2.582671
       vertex -2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552208
    outer loop
       vertex -0.000000 5.543277 -2.296101
       vertex 0.878680 5.543277 -2.121320
       vertex 1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552208
    outer loop
       vertex -0.000000 5.543277 -2.296101
       vertex 1.189829 5.000000 -2.872500
       vertex 0.678174 5.000000 -2.974275
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552209
    outer loop
       vertex -0.000000 5.543277 -2.296101
       vertex 0.678174 5.000000 -2.974275
       vertex -0.000000 5.000000 -3.109172
    endloop
  endfacet
  facet normal 0.552208 0.826439 0.109841
    outer loop
       vertex 2.296101 5.543277 0.000000
       vertex 2.121320 5.543277 0.878680
       vertex 2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal 0.552208 0.826439 0.109841
    outer loop
       vertex 2.296101 5.543277 0.000000
       vertex 2.872500 5.000000 1.189829
       vertex 2.974275 5.000000 0.678174
    endloop
  endfacet
  facet normal 0.552209 0.826439 0.109841
    outer loop
       vertex 2.296101 5.543277 0.000000
       vertex 2.974275 5.000000 0.678174
       vertex 3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal -0.468140 0.826439 -0.312801
    outer loop
       vertex -2.121320 5.543277 -0.878680
       vertex -1.623588 5.543277 -1.623588
       vertex -2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.468140 0.826438 -0.312801
    outer loop
       vertex -2.121320 5.543277 -0.878680
       vertex -2.198517 5.000000 -2.198517
       vertex -2.488345 5.000000 -1.764757
    endloop
  endfacet
  facet normal -0.468140 0.826439 -0.312801
    outer loop
       vertex -2.121320 5.543277 -0.878680
       vertex -2.488345 5.000000 -1.764757
       vertex -2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.194944 0.980048 0.038777
    outer loop
       vertex 0.000000 6.000000 0.000000
       vertex 2.121320 5.543277 0.878680
       vertex 2.296101 5.543277 0.000000
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex -0.878680 5.543277 2.121320
       vertex -1.623588 5.543277 1.623588
       vertex -2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal -0.312801 0.826438 0.468140
    outer loop
       vertex -0.878680 5.543277 2.121320
       vertex -2.198517 5.000000 2.198517
       vertex -1.764757 5.000000 2.488345
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex -0.878680 5.543277 2.121320
       vertex -1.764757 5.000000 2.488345
       vertex -1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex 0.878680 5.543277 -2.121320
       vertex 1.623588 5.543277 -1.623588
       vertex 2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.312801 0.826438 -0.468140
    outer loop
       vertex 0.878680 5.543277 -2.121320
       vertex 2.198517 5.000000 -2.198517
       vertex 1.764757 5.000000 -2.488345
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex 0.878680 5.543277 -2.121320
       vertex 1.764757 5.000000 -2.488345
       vertex 1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.110427 0.980048 0.165265
    outer loop
       vertex -0.000000 6.000000 0.000000
       vertex -1.623588 5.543277 1.623588
       vertex -0.878680 5.543277 2.121320
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex 1.623588 5.543277 -1.623588
       vertex 2.121320 5.543277 -0.878680
       vertex 2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex 1.623588 5.543277 -1.623588
       vertex 2.872500 5.000000 -1.189829
       vertex 2.582671 5.000000 -1.623588
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex 1.623588 5.543277 -1.623588
       vertex 2.582671 5.000000 -1.623588
       vertex 2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.552208 0.826439 -0.109841
    outer loop
       vertex -2.296101 5.543277 0.000000
       vertex -2.121320 5.543277 -0.878680
       vertex -2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.552208 0.826439 -0.109841
    outer loop
       vertex -2.296101 5.543277 0.000000
       vertex -2.872500 5.000000 -1.189829
       vertex -2.974275 5.000000 -0.678174
    endloop
  endfacet
  facet normal -0.552209 0.826439 -0.109841
    outer loop
       vertex -2.296101 5.543277 0.000000
       vertex -2.974275 5.000000 -0.678174
       vertex -3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552209
    outer loop
       vertex 0.878680 5.543277 2.121320
       vertex 0.000000 5.543277 2.296101
       vertex 0.000000 5.000000 3.109172
    endloop
  endfacet
  facet normal 0.109841 0.826438 0.552209
    outer loop
       vertex 0.878680 5.543277 2.121320
       vertex 0.000000 5.000000 3.109172
       vertex 0.511654 5.000000 3.007397
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552208
    outer loop
       vertex 0.878680 5.543277 2.121320
       vertex 0.511654 5.000000 3.007397
       vertex 1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552209
    outer loop
       vertex -0.878680 5.543277 -2.121320
       vertex -0.000000 5.543277 -2.296101
       vertex -0.000000 5.000000 -3.109172
    endloop
  endfacet
  facet normal -0.109841 0.826438 -0.552209
    outer loop
       vertex -0.878680 5.543277 -2.121320
       vertex -0.000000 5.000000 -3.109172
       vertex -0.511654 5.000000 -3.007397
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552208
    outer loop
       vertex -0.878680 5.543277 -2.121320
       vertex -0.511654 5.000000 -3.007397
       vertex -1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.038777 0.980048 0.194944
    outer loop
       vertex 0.000000 6.000000 0.000000
       vertex -0.878680 5.543277 2.121320
       vertex 0.000000 5.543277 2.296101
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex -1.623588 5.543277 1.623588
       vertex -2.121320 5.543277 0.878680
       vertex -2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex -1.623588 5.543277 1.623588
       vertex -2.872500 5.000000 1.189829
       vertex -2.582671 5.000000 1.623588
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex -1.623588 5.543277 1.623588
       vertex -2.582671 5.000000 1.623588
       vertex -2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552208
    outer loop
       vertex 0.000000 5.543277 2.296101
       vertex -0.878680 5.543277 2.121320
       vertex -1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552208
    outer loop
       vertex 0.000000 5.543277 2.296101
       vertex -1.189829 5.000000 2.872500
       vertex -0.678174 5.000000 2.974275
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552209
    outer loop
       vertex 0.000000 5.543277 2.296101
       vertex -0.678174 5.000000 2.974275
       vertex 0.000000 5.000000 3.109172
    endloop
  endfacet
  facet normal 0.194944 0.980048 -0.038777
    outer loop
       vertex 0.000000 6.000000 -0.000000
       vertex 2.296101 5.543277 -0.000000
       vertex 2.121320 5.543277 -0.878680
    endloop
  endfacet
  facet normal -0.165265 0.980048 -0.110427
    outer loop
       vertex -0.000000 6.000000 -0.000000
       vertex -1.623588 5.543277 -1.623588
       vertex -2.121320 5.543277 -0.878680
    endloop
  endfacet
  facet normal -0.110427 0.980048 -0.165265
    outer loop
       vertex -0.000000 6.000000 -0.000000
       vertex -0.878680 5.543277 -2.121320
       vertex -1.623588 5.543277 -1.623588
    endloop
  endfacet
  facet normal 0.110427 0.980048 0.165265
    outer loop
       vertex 0.000000 6.000000 0.000000
       vertex 0.878680 5.543277 2.121320
       vertex 1.623588 5.543277 1.623588
    endloop
  endfacet
  facet normal 0.165265 0.980048 -0.110427
    outer loop
       vertex 0.000000 6.000000 -0.000000
       vertex 2.121320 5.543277 -0.878680
       vertex 1.623588 5.543277 -1.623588
    endloop
  endfacet
  facet normal 0.038777 0.980048 0.194944
    outer loop
       vertex 0.000000 6.000000 0.000000
       vertex 0.000000 5.543277 2.296101
       vertex 0.878680 5.543277 2.121320
    endloop
  endfacet
  facet normal -0.194944 0.980048 0.038777
    outer loop
       vertex -0.000000 6.000000 0.000000
       vertex -2.296101 5.543277 0.000000
       vertex -2.121320 5.543277 0.878680
    endloop
  endfacet
  facet normal 0.165265 0.980048 0.110427
    outer loop
       vertex 0.000000 6.000000 0.000000
       vertex 1.623588 5.543277 1.623588
       vertex 2.121320 5.543277 0.878680
    endloop
  endfacet
  facet normal -0.194944 0.980048 -0.038777
    outer loop
       vertex -0.000000 6.000000 0.000000
       vertex -2.121320 5.543277 -0.878680
       vertex -2.296101 5.543277 0.000000
    endloop
  endfacet
  facet normal 0.110427 0.980048 -0.165265
    outer loop
       vertex 0.000000 6.000000 -0.000000
       vertex 1.623588 5.543277 -1.623588
       vertex 0.878680 5.543277 -2.121320
    endloop
  endfacet
  facet normal 0.038777 0.980048 -0.194944
    outer loop
       vertex -0.000000 6.000000 -0.000000
       vertex 0.878680 5.543277 -2.121320
       vertex -0.000000 5.543277 -2.296101
    endloop
  endfacet
  facet normal -0.038777 0.980048 -0.194944
    outer loop
       vertex -0.000000 6.000000 -0.000000
       vertex -0.000000 5.543277 -2.296101
       vertex -0.878680 5.543277 -2.121320
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex 5.121321 -2.296101 -2.121320
       vertex 5.543277 -2.296101 -0.000000
       vertex 5.000000 -3.109172 -0.000000
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex 5.121321 -2.296101 -2.121320
       vertex 5.000000 -3.109172 -0.000000
       vertex 5.000000 -2.564862 -1.828426
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163172
    outer loop
       vertex 5.121321 -2.296101 -2.121320
       vertex 5.000000 -2.564862 -1.828426
       vertex 5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex 5.000000 2.296101 -2.302890
       vertex 5.121321 2.296101 -2.121320
       vertex 5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex 5.000000 2.296101 -2.302890
       vertex 5.543277 0.000000 -2.296101
       vertex 5.000000 0.768310 -2.839378
    endloop
  endfacet
  facet normal 0.816084 0.191480 -0.545290
    outer loop
       vertex 5.000000 0.768310 -2.839378
       vertex 5.543277 0.000000 -2.296101
       vertex 5.000000 0.000000 -3.109172
    endloop
  endfacet
  facet normal 0.962637 -0.191480 0.191481
    outer loop
       vertex 6.000000 0.000000 0.000000
       vertex 5.543277 0.000000 2.296101
       vertex 5.121321 -2.296101 2.121320
    endloop
  endfacet
  facet normal 0.962637 -0.191481 0.191480
    outer loop
       vertex 6.000000 0.000000 0.000000
       vertex 5.121321 -2.296101 2.121320
       vertex 5.543277 -2.296101 0.000000
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex 5.121321 2.296101 2.121320
       vertex 5.000000 2.296101 2.302890
       vertex 5.000000 1.979074 2.414214
    endloop
  endfacet
  facet normal 0.816084 0.191480 0.545290
    outer loop
       vertex 5.121321 2.296101 2.121320
       vertex 5.000000 1.979074 2.414214
       vertex 5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex 5.121321 2.296101 2.121320
       vertex 5.000000 0.000000 3.109172
       vertex 5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal 0.816084 -0.191480 -0.545290
    outer loop
       vertex 5.000000 0.000000 -3.109172
       vertex 5.543277 0.000000 -2.296101
       vertex 5.121321 -2.296101 -2.121320
    endloop
  endfacet
  facet normal 0.816084 -0.191480 -0.545290
    outer loop
       vertex 5.000000 0.000000 -3.109172
       vertex 5.121321 -2.296101 -2.121320
       vertex 5.000000 -1.979074 -2.414214
    endloop
  endfacet
  facet normal 0.816085 -0.191480 -0.545290
    outer loop
       vertex 5.000000 -1.979074 -2.414214
       vertex 5.121321 -2.296101 -2.121320
       vertex 5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464677
    outer loop
       vertex 5.000000 -2.296101 -2.302890
       vertex 5.121321 -2.296101 -2.121320
       vertex 5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex 5.000000 2.296101 2.302890
       vertex 5.121321 2.296101 2.121320
       vertex 5.000000 2.492630 2.071068
    endloop
  endfacet
  facet normal 0.695438 0.548124 -0.464677
    outer loop
       vertex 5.000000 2.492630 -2.071068
       vertex 5.121321 2.296101 -2.121320
       vertex 5.000000 2.407425 -2.171573
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex 5.000000 2.407425 -2.171573
       vertex 5.121321 2.296101 -2.121320
       vertex 5.000000 2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex 5.543277 -2.296101 0.000000
       vertex 5.121321 -2.296101 2.121320
       vertex 5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex 5.543277 -2.296101 0.000000
       vertex 5.000000 -2.492630 2.071068
       vertex 5.000000 -2.947442 0.543277
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex 5.543277 -2.296101 0.000000
       vertex 5.000000 -2.947442 0.543277
       vertex 5.000000 -3.109172 0.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex 5.000000 3.109172 -0.000000
       vertex 5.543277 2.296101 -0.000000
       vertex 5.000000 2.947442 -0.543277
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex 5.000000 2.947442 -0.543277
       vertex 5.543277 2.296101 -0.000000
       vertex 5.121321 2.296101 -2.121320
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex 5.000000 2.947442 -0.543277
       vertex 5.121321 2.296101 -2.121320
       vertex 5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163172
    outer loop
       vertex 5.000000 2.492630 2.071068
       vertex 5.121321 2.296101 2.121320
       vertex 5.000000 2.564862 1.828426
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex 5.000000 2.564862 1.828426
       vertex 5.121321 2.296101 2.121320
       vertex 5.543277 2.296101 0.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex 5.000000 2.564862 1.828426
       vertex 5.543277 2.296101 0.000000
       vertex 5.000000 3.109172 0.000000
    endloop
  endfacet
  facet normal 0.816084 -0.191480 0.545290
    outer loop
       vertex 5.543277 0.000000 2.296101
       vertex 5.000000 0.000000 3.109172
       vertex 5.000000 -0.768310 2.839378
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex 5.543277 0.000000 2.296101
       vertex 5.000000 -0.768310 2.839378
       vertex 5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex 5.543277 0.000000 2.296101
       vertex 5.000000 -2.296101 2.302890
       vertex 5.121321 -2.296101 2.121320
    endloop
  endfacet
  facet normal 0.962637 0.191480 0.191480
    outer loop
       vertex 5.543277 2.296101 0.000000
       vertex 5.121321 2.296101 2.121320
       vertex 5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal 0.962637 0.191481 0.191481
    outer loop
       vertex 5.543277 2.296101 0.000000
       vertex 5.543277 0.000000 2.296101
       vertex 6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex 5.121321 -2.296101 2.121320
       vertex 5.000000 -2.296101 2.302890
       vertex 5.000000 -2.407425 2.171573
    endloop
  endfacet
  facet normal 0.695438 -0.548124 0.464678
    outer loop
       vertex 5.121321 -2.296101 2.121320
       vertex 5.000000 -2.407425 2.171573
       vertex 5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal 0.962637 0.191481 -0.191480
    outer loop
       vertex 5.121321 2.296101 -2.121320
       vertex 5.543277 2.296101 -0.000000
       vertex 6.000000 0.000000 -0.000000
    endloop
  endfacet
  facet normal 0.962637 0.191480 -0.191481
    outer loop
       vertex 5.121321 2.296101 -2.121320
       vertex 6.000000 0.000000 -0.000000
       vertex 5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal 0.962637 -0.191481 -0.191481
    outer loop
       vertex 5.543277 0.000000 -2.296101
       vertex 6.000000 0.000000 -0.000000
       vertex 5.543277 -2.296101 -0.000000
    endloop
  endfacet
  facet normal 0.962637 -0.191480 -0.191480
    outer loop
       vertex 5.543277 0.000000 -2.296101
       vertex 5.543277 -2.296101 -0.000000
       vertex 5.121321 -2.296101 -2.121320
    endloop
  endfacet
  facet normal -0.191481 -0.191480 0.962637
    outer loop
       vertex 0.000000 0.000000 6.000000
       vertex -2.296101 0.000000 5.543277
       vertex -2.121320 -2.296101 5.121321
    endloop
  endfacet
  facet normal -0.191480 -0.191481 0.962637
    outer loop
       vertex 0.000000 0.000000 6.000000
       vertex -2.121320 -2.296101 5.121321
       vertex 0.000000 -2.296101 5.543277
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816084
    outer loop
       vertex 3.109172 0.000000 5.000000
       vertex 2.296101 0.000000 5.543277
       vertex 2.121320 -2.296101 5.121321
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816084
    outer loop
       vertex 3.109172 0.000000 5.000000
       vertex 2.121320 -2.296101 5.121321
       vertex 2.414214 -1.979074 5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816085
    outer loop
       vertex 2.414214 -1.979074 5.000000
       vertex 2.121320 -2.296101 5.121321
       vertex 2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816084
    outer loop
       vertex -2.296101 0.000000 5.543277
       vertex -3.109172 0.000000 5.000000
       vertex -2.839378 -0.768310 5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex -2.296101 0.000000 5.543277
       vertex -2.839378 -0.768310 5.000000
       vertex -2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex -2.296101 0.000000 5.543277
       vertex -2.302890 -2.296101 5.000000
       vertex -2.121320 -2.296101 5.121321
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex -2.121320 -2.296101 5.121321
       vertex -2.302890 -2.296101 5.000000
       vertex -2.171573 -2.407425 5.000000
    endloop
  endfacet
  facet normal -0.464678 -0.548124 0.695438
    outer loop
       vertex -2.121320 -2.296101 5.121321
       vertex -2.171573 -2.407425 5.000000
       vertex -2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal -0.191480 0.191480 0.962637
    outer loop
       vertex 0.000000 2.296101 5.543277
       vertex -2.121320 2.296101 5.121321
       vertex -2.296101 0.000000 5.543277
    endloop
  endfacet
  facet normal -0.191481 0.191481 0.962637
    outer loop
       vertex 0.000000 2.296101 5.543277
       vertex -2.296101 0.000000 5.543277
       vertex 0.000000 0.000000 6.000000
    endloop
  endfacet
  facet normal 0.191481 -0.191481 0.962637
    outer loop
       vertex 2.296101 0.000000 5.543277
       vertex 0.000000 0.000000 6.000000
       vertex 0.000000 -2.296101 5.543277
    endloop
  endfacet
  facet normal 0.191480 -0.191480 0.962637
    outer loop
       vertex 2.296101 0.000000 5.543277
       vertex 0.000000 -2.296101 5.543277
       vertex 2.121320 -2.296101 5.121321
    endloop
  endfacet
  facet normal 0.191480 0.191481 0.962637
    outer loop
       vertex 2.121320 2.296101 5.121321
       vertex 0.000000 2.296101 5.543277
       vertex 0.000000 0.000000 6.000000
    endloop
  endfacet
  facet normal 0.191481 0.191480 0.962637
    outer loop
       vertex 2.121320 2.296101 5.121321
       vertex 0.000000 0.000000 6.000000
       vertex 2.296101 0.000000 5.543277
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex -2.121320 2.296101 5.121321
       vertex -2.302890 2.296101 5.000000
       vertex -2.414214 1.979074 5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816084
    outer loop
       vertex -2.121320 2.296101 5.121321
       vertex -2.414214 1.979074 5.000000
       vertex -3.109172 0.000000 5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex -2.121320 2.296101 5.121321
       vertex -3.109172 0.000000 5.000000
       vertex -2.296101 0.000000 5.543277
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex 2.302890 2.296101 5.000000
       vertex 2.121320 2.296101 5.121321
       vertex 2.296101 0.000000 5.543277
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex 2.302890 2.296101 5.000000
       vertex 2.296101 0.000000 5.543277
       vertex 2.839378 0.768310 5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816084
    outer loop
       vertex 2.839378 0.768310 5.000000
       vertex 2.296101 0.000000 5.543277
       vertex 3.109172 0.000000 5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex 0.000000 3.109172 5.000000
       vertex 0.000000 2.296101 5.543277
       vertex 0.543277 2.947442 5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex 0.543277 2.947442 5.000000
       vertex 0.000000 2.296101 5.543277
       vertex 2.121320 2.296101 5.121321
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex 0.543277 2.947442 5.000000
       vertex 2.121320 2.296101 5.121321
       vertex 2.071068 2.492630 5.000000
    endloop
  endfacet
  facet normal -0.163172 0.548124 0.820326
    outer loop
       vertex -2.071068 2.492630 5.000000
       vertex -2.121320 2.296101 5.121321
       vertex -1.828426 2.564862 5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex -1.828426 2.564862 5.000000
       vertex -2.121320 2.296101 5.121321
       vertex 0.000000 2.296101 5.543277
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex -1.828426 2.564862 5.000000
       vertex 0.000000 2.296101 5.543277
       vertex 0.000000 3.109172 5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695438
    outer loop
       vertex 2.071068 2.492630 5.000000
       vertex 2.121320 2.296101 5.121321
       vertex 2.171573 2.407425 5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex 2.171573 2.407425 5.000000
       vertex 2.121320 2.296101 5.121321
       vertex 2.302890 2.296101 5.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 0.695439
    outer loop
       vertex 2.302890 -2.296101 5.000000
       vertex 2.121320 -2.296101 5.121321
       vertex 2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex -2.302890 2.296101 5.000000
       vertex -2.121320 2.296101 5.121321
       vertex -2.071068 2.492630 5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex 2.121320 -2.296101 5.121321
       vertex 0.000000 -2.296101 5.543277
       vertex 0.000000 -3.109172 5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex 2.121320 -2.296101 5.121321
       vertex 0.000000 -3.109172 5.000000
       vertex 1.828426 -2.564862 5.000000
    endloop
  endfacet
  facet normal 0.163172 -0.548124 0.820326
    outer loop
       vertex 2.121320 -2.296101 5.121321
       vertex 1.828426 -2.564862 5.000000
       vertex 2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex 0.000000 -2.296101 5.543277
       vertex -2.121320 -2.296101 5.121321
       vertex -2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex 0.000000 -2.296101 5.543277
       vertex -2.071068 -2.492630 5.000000
       vertex -0.543277 -2.947442 5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex 0.000000 -2.296101 5.543277
       vertex -0.543277 -2.947442 5.000000
       vertex 0.000000 -3.109172 5.000000
    endloop
  endfacet
endsolid
```

### Intersection

```stl
solid ScadSharp
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071068
       vertex 5.000000 3.109171 0.000001
       vertex 5.000000 2.492629 2.071070
    endloop
  endfacet
  facet normal 1.000000 0.000000 -0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071068
       vertex 5.000000 2.492629 2.071070
       vertex 5.000000 2.296102 2.302890
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071068
       vertex 5.000000 2.296102 2.302890
       vertex 5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071068
       vertex 5.000000 0.000000 3.109172
       vertex 5.000000 -2.296099 2.302891
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071068
       vertex 5.000000 -2.296099 2.302891
       vertex 5.000000 -2.299215 2.299215
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 2.492629 -2.071068
       vertex 5.000000 -2.299215 2.299215
       vertex 5.000000 2.299215 -2.299215
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492630 2.071066
       vertex 5.000000 -3.109172 0.000000
       vertex 5.000000 -2.492629 -2.071069
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492630 2.071066
       vertex 5.000000 -2.492629 -2.071069
       vertex 5.000000 -2.296100 -2.302891
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492630 2.071066
       vertex 5.000000 -2.296100 -2.302891
       vertex 5.000000 0.000001 -3.109173
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492630 2.071066
       vertex 5.000000 0.000001 -3.109173
       vertex 5.000000 2.296099 -2.302891
    endloop
  endfacet
  facet normal 1.000000 0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492630 2.071066
       vertex 5.000000 2.296099 -2.302891
       vertex 5.000000 2.299215 -2.299215
    endloop
  endfacet
  facet normal 1.000000 -0.000000 0.000000
    outer loop
       vertex 5.000000 -2.492630 2.071066
       vertex 5.000000 2.299215 -2.299215
       vertex 5.000000 -2.299215 2.299215
    endloop
  endfacet
  facet normal 0.000000 -0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex 3.109172 -0.000001 5.000000
       vertex 2.302890 2.296101 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex 2.302890 2.296101 5.000000
       vertex 2.071070 2.492628 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex 2.071070 2.492628 5.000000
       vertex 0.000000 3.109171 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex 0.000000 3.109171 5.000000
       vertex -2.071069 2.492628 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex -2.071069 2.492628 5.000000
       vertex -2.299215 2.299215 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex 2.302890 -2.296100 5.000000
       vertex -2.299215 2.299215 5.000000
       vertex 2.299215 -2.299215 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302890 2.296100 5.000000
       vertex -3.109172 -0.000001 5.000000
       vertex -2.302891 -2.296098 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302890 2.296100 5.000000
       vertex -2.302891 -2.296098 5.000000
       vertex -2.071067 -2.492630 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302890 2.296100 5.000000
       vertex -2.071067 -2.492630 5.000000
       vertex 0.000002 -3.109172 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302890 2.296100 5.000000
       vertex 0.000002 -3.109172 5.000000
       vertex 2.071066 -2.492631 5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 1.000000
    outer loop
       vertex -2.302890 2.296100 5.000000
       vertex 2.071066 -2.492631 5.000000
       vertex 2.299215 -2.299215 5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 1.000000
    outer loop
       vertex -2.302890 2.296100 5.000000
       vertex 2.299215 -2.299215 5.000000
       vertex -2.299215 2.299215 5.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 0.000003 5.000000 3.109172
       vertex 1.189825 5.000000 2.872501
    endloop
  endfacet
  facet normal -0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 1.189825 5.000000 2.872501
       vertex 2.198517 5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 2.198517 5.000000 2.198516
       vertex 2.872499 5.000000 1.189830
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 2.872499 5.000000 1.189830
       vertex 3.109172 5.000000 -0.000004
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 3.109172 5.000000 -0.000004
       vertex 2.872501 5.000000 -1.189825
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 2.872501 5.000000 -1.189825
       vertex 2.198516 5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex 2.198516 5.000000 -2.198516
       vertex -2.198516 5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex 0.000001 5.000000 -3.109172
       vertex -1.189830 5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex -1.189830 5.000000 -2.872500
       vertex -2.198516 5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex -2.198516 5.000000 -2.198517
       vertex -2.872501 5.000000 -1.189827
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex -2.872501 5.000000 -1.189827
       vertex -3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex -3.109172 5.000000 0.000000
       vertex -2.872502 5.000000 1.189824
    endloop
  endfacet
  facet normal 0.000000 1.000000 0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex -2.872502 5.000000 1.189824
       vertex -2.198516 5.000000 2.198516
    endloop
  endfacet
  facet normal 0.000000 1.000000 -0.000000
    outer loop
       vertex 1.189830 5.000000 -2.872499
       vertex -2.198516 5.000000 2.198516
       vertex 2.198516 5.000000 -2.198516
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -0.000000 -5.000000 3.109172
       vertex -1.189827 -5.000000 2.872501
    endloop
  endfacet
  facet normal -0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -1.189827 -5.000000 2.872501
       vertex -2.198517 -5.000000 2.198515
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -2.198517 -5.000000 2.198515
       vertex -2.872499 -5.000000 1.189829
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -2.872499 -5.000000 1.189829
       vertex -3.109171 -5.000000 -0.000000
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -3.109171 -5.000000 -0.000000
       vertex -2.872500 -5.000000 -1.189828
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -2.872500 -5.000000 -1.189828
       vertex -2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.000000 -1.000000 -0.000000
    outer loop
       vertex 1.189828 -5.000000 2.872501
       vertex -2.198517 -5.000000 -2.198517
       vertex 2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 0.000000 -5.000000 -3.109172
       vertex 1.189826 -5.000000 -2.872501
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 1.189826 -5.000000 -2.872501
       vertex 2.198514 -5.000000 -2.198518
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 2.198514 -5.000000 -2.198518
       vertex 2.872500 -5.000000 -1.189828
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 2.872500 -5.000000 -1.189828
       vertex 3.109172 -5.000000 0.000004
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 3.109172 -5.000000 0.000004
       vertex 2.872502 -5.000000 1.189827
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 2.872502 -5.000000 1.189827
       vertex 2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal 0.000000 -1.000000 0.000000
    outer loop
       vertex -1.189826 -5.000000 -2.872501
       vertex 2.198517 -5.000000 2.198517
       vertex -2.198516 -5.000000 -2.198516
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071070
       vertex -5.000000 3.109171 0.000001
       vertex -5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal -1.000000 0.000000 -0.000000
    outer loop
       vertex -5.000000 2.492629 2.071070
       vertex -5.000000 2.492630 -2.071068
       vertex -5.000000 2.296102 -2.302888
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071070
       vertex -5.000000 2.296102 -2.302888
       vertex -5.000000 -0.000001 -3.109172
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071070
       vertex -5.000000 -0.000001 -3.109172
       vertex -5.000000 -2.296099 -2.302891
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071070
       vertex -5.000000 -2.296099 -2.302891
       vertex -5.000000 -2.299215 -2.299215
    endloop
  endfacet
  facet normal -1.000000 -0.000000 0.000000
    outer loop
       vertex -5.000000 2.492629 2.071070
       vertex -5.000000 -2.299215 -2.299215
       vertex -5.000000 2.299216 2.299216
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071068
       vertex -5.000000 -3.109171 0.000000
       vertex -5.000000 -2.492629 2.071068
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071068
       vertex -5.000000 -2.492629 2.071068
       vertex -5.000000 -2.296099 2.302891
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071068
       vertex -5.000000 -2.296099 2.302891
       vertex -5.000000 0.000000 3.109173
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071068
       vertex -5.000000 0.000000 3.109173
       vertex -5.000000 2.296099 2.302891
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071068
       vertex -5.000000 2.296099 2.302891
       vertex -5.000000 2.299215 2.299215
    endloop
  endfacet
  facet normal -1.000000 0.000000 0.000000
    outer loop
       vertex -5.000000 -2.492629 -2.071068
       vertex -5.000000 2.299215 2.299215
       vertex -5.000000 -2.299215 -2.299215
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302887 2.296103 -5.000000
       vertex 3.109172 -0.000002 -5.000000
       vertex 2.302891 -2.296099 -5.000000
    endloop
  endfacet
  facet normal 0.000000 -0.000000 -1.000000
    outer loop
       vertex 2.302887 2.296103 -5.000000
       vertex 2.302891 -2.296099 -5.000000
       vertex 2.071067 -2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302887 2.296103 -5.000000
       vertex 2.071067 -2.492630 -5.000000
       vertex -0.000001 -3.109172 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302887 2.296103 -5.000000
       vertex -0.000001 -3.109172 -5.000000
       vertex -2.071069 -2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302887 2.296103 -5.000000
       vertex -2.071069 -2.492630 -5.000000
       vertex -2.299216 -2.299216 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.000000 -1.000000
    outer loop
       vertex 2.302887 2.296103 -5.000000
       vertex -2.299216 -2.299216 -5.000000
       vertex 2.299215 2.299215 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex -3.109172 -0.000000 -5.000000
       vertex -2.302890 2.296099 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex -2.302890 2.296099 -5.000000
       vertex -2.071067 2.492629 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex -2.071067 2.492629 -5.000000
       vertex 0.000000 3.109171 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex 0.000000 3.109171 -5.000000
       vertex 2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex 2.071068 2.492630 -5.000000
       vertex 2.299215 2.299215 -5.000000
    endloop
  endfacet
  facet normal 0.000000 0.000000 -1.000000
    outer loop
       vertex -2.302889 -2.296103 -5.000000
       vertex 2.299215 2.299215 -5.000000
       vertex -2.299216 -2.299216 -5.000000
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex 1.189829 5.000000 2.872500
       vertex 1.623588 4.242640 3.919689
       vertex 1.623588 5.000000 2.582671
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 5.000000 2.582671
       vertex 1.623588 4.242640 3.919689
       vertex 3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal 0.312801 0.826439 0.468140
    outer loop
       vertex 1.623588 5.000000 2.582671
       vertex 3.000000 4.242640 3.000000
       vertex 2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552209
    outer loop
       vertex -1.623588 -4.242640 -3.919689
       vertex -0.000000 -4.242640 -4.242640
       vertex -0.000000 -5.000000 -3.109172
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552209
    outer loop
       vertex -1.623588 -4.242640 -3.919689
       vertex -0.000000 -5.000000 -3.109172
       vertex -0.678174 -5.000000 -2.974275
    endloop
  endfacet
  facet normal -0.109841 -0.826439 -0.552208
    outer loop
       vertex -1.623588 -4.242640 -3.919689
       vertex -0.678174 -5.000000 -2.974275
       vertex -1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.468140 0.826438 -0.312801
    outer loop
       vertex -2.198517 5.000000 -2.198517
       vertex -3.000000 4.242640 -3.000000
       vertex -2.488345 5.000000 -1.764757
    endloop
  endfacet
  facet normal -0.468140 0.826439 -0.312801
    outer loop
       vertex -2.488345 5.000000 -1.764757
       vertex -3.000000 4.242640 -3.000000
       vertex -3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.468140 0.826438 -0.312801
    outer loop
       vertex -2.488345 5.000000 -1.764757
       vertex -3.919689 4.242640 -1.623588
       vertex -2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464677
    outer loop
       vertex -3.919689 -2.296101 3.919689
       vertex -5.000000 -2.296101 2.302890
       vertex -5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464677
    outer loop
       vertex -3.919689 -2.296101 3.919689
       vertex -5.000000 -2.492630 2.071068
       vertex -3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal -0.695439 -0.548124 0.464677
    outer loop
       vertex -3.919689 -2.296101 3.919689
       vertex -3.919689 -4.242640 1.623588
       vertex -3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex 3.919689 4.242640 1.623588
       vertex 3.000000 4.242640 3.000000
       vertex 3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex 3.919689 4.242640 1.623588
       vertex 3.919689 2.296101 3.919689
       vertex 5.000000 2.296101 2.302890
    endloop
  endfacet
  facet normal 0.695439 0.548124 0.464677
    outer loop
       vertex 3.919689 4.242640 1.623588
       vertex 5.000000 2.296101 2.302890
       vertex 5.000000 2.492630 2.071068
    endloop
  endfacet
  facet normal 0.552208 0.826439 0.109841
    outer loop
       vertex 2.872500 5.000000 1.189829
       vertex 3.919689 4.242640 1.623588
       vertex 2.974275 5.000000 0.678174
    endloop
  endfacet
  facet normal 0.552209 0.826439 0.109841
    outer loop
       vertex 2.974275 5.000000 0.678174
       vertex 3.919689 4.242640 1.623588
       vertex 4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal 0.552209 0.826439 0.109841
    outer loop
       vertex 2.974275 5.000000 0.678174
       vertex 4.242640 4.242640 0.000000
       vertex 3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal 0.468140 0.826438 0.312801
    outer loop
       vertex 2.198517 5.000000 2.198517
       vertex 3.000000 4.242640 3.000000
       vertex 2.488345 5.000000 1.764757
    endloop
  endfacet
  facet normal 0.468140 0.826439 0.312801
    outer loop
       vertex 2.488345 5.000000 1.764757
       vertex 3.000000 4.242640 3.000000
       vertex 3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal 0.468140 0.826438 0.312801
    outer loop
       vertex 2.488345 5.000000 1.764757
       vertex 3.919689 4.242640 1.623588
       vertex 2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal 0.816084 -0.191480 0.545290
    outer loop
       vertex 5.000000 0.000000 3.109172
       vertex 4.242640 0.000000 4.242640
       vertex 3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex 5.000000 0.000000 3.109172
       vertex 3.919689 -2.296101 3.919689
       vertex 5.000000 -0.768310 2.839378
    endloop
  endfacet
  facet normal 0.816085 -0.191480 0.545290
    outer loop
       vertex 5.000000 -0.768310 2.839378
       vertex 3.919689 -2.296101 3.919689
       vertex 5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex -3.000000 4.242640 -3.000000
       vertex -1.623588 4.242640 -3.919689
       vertex -2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex -3.000000 4.242640 -3.000000
       vertex -2.071068 2.492630 -5.000000
       vertex -2.171573 2.407425 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex -3.000000 4.242640 -3.000000
       vertex -2.171573 2.407425 -5.000000
       vertex -2.302890 2.296101 -5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 -0.695439
    outer loop
       vertex -3.000000 4.242640 -3.000000
       vertex -2.302890 2.296101 -5.000000
       vertex -3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex -3.000000 -4.242640 3.000000
       vertex -3.919689 -4.242640 1.623588
       vertex -2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal -0.468140 -0.826439 0.312801
    outer loop
       vertex -3.000000 -4.242640 3.000000
       vertex -2.872500 -5.000000 1.189829
       vertex -2.488345 -5.000000 1.764757
    endloop
  endfacet
  facet normal -0.468140 -0.826438 0.312801
    outer loop
       vertex -3.000000 -4.242640 3.000000
       vertex -2.488345 -5.000000 1.764757
       vertex -2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552209
    outer loop
       vertex -0.000000 5.000000 -3.109172
       vertex -0.000000 4.242640 -4.242640
       vertex -0.511654 5.000000 -3.007397
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552209
    outer loop
       vertex -0.511654 5.000000 -3.007397
       vertex -0.000000 4.242640 -4.242640
       vertex -1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.109841 0.826439 -0.552209
    outer loop
       vertex -0.511654 5.000000 -3.007397
       vertex -1.623588 4.242640 -3.919689
       vertex -1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.816085 -0.191480 -0.545290
    outer loop
       vertex 4.242640 0.000000 -4.242640
       vertex 5.000000 0.000000 -3.109172
       vertex 5.000000 -1.979074 -2.414214
    endloop
  endfacet
  facet normal 0.816085 -0.191480 -0.545290
    outer loop
       vertex 4.242640 0.000000 -4.242640
       vertex 5.000000 -1.979074 -2.414214
       vertex 5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.816085 -0.191480 -0.545290
    outer loop
       vertex 4.242640 0.000000 -4.242640
       vertex 5.000000 -2.296101 -2.302890
       vertex 3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex -1.189829 5.000000 -2.872500
       vertex -1.623588 4.242640 -3.919689
       vertex -1.623588 5.000000 -2.582671
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 5.000000 -2.582671
       vertex -1.623588 4.242640 -3.919689
       vertex -3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.312801 0.826439 -0.468140
    outer loop
       vertex -1.623588 5.000000 -2.582671
       vertex -3.000000 4.242640 -3.000000
       vertex -2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex 3.919689 4.242640 -1.623588
       vertex 4.242640 4.242640 -0.000000
       vertex 5.000000 3.109172 -0.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex 3.919689 4.242640 -1.623588
       vertex 5.000000 3.109172 -0.000000
       vertex 5.000000 2.947442 -0.543277
    endloop
  endfacet
  facet normal 0.820326 0.548124 -0.163173
    outer loop
       vertex 3.919689 4.242640 -1.623588
       vertex 5.000000 2.947442 -0.543277
       vertex 5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex 3.919689 -4.242640 1.623588
       vertex 3.000000 -4.242640 3.000000
       vertex 2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex 3.919689 -4.242640 1.623588
       vertex 2.198517 -5.000000 2.198517
       vertex 2.582671 -5.000000 1.623588
    endloop
  endfacet
  facet normal 0.468140 -0.826439 0.312801
    outer loop
       vertex 3.919689 -4.242640 1.623588
       vertex 2.582671 -5.000000 1.623588
       vertex 2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex 5.000000 2.296101 2.302890
       vertex 3.919689 2.296101 3.919689
       vertex 4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex 5.000000 2.296101 2.302890
       vertex 4.242640 0.000000 4.242640
       vertex 5.000000 1.979074 2.414214
    endloop
  endfacet
  facet normal 0.816085 0.191480 0.545290
    outer loop
       vertex 5.000000 1.979074 2.414214
       vertex 4.242640 0.000000 4.242640
       vertex 5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex -3.919689 -4.242640 -1.623588
       vertex -3.000000 -4.242640 -3.000000
       vertex -2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex -3.919689 -4.242640 -1.623588
       vertex -2.198517 -5.000000 -2.198517
       vertex -2.582671 -5.000000 -1.623588
    endloop
  endfacet
  facet normal -0.468140 -0.826439 -0.312801
    outer loop
       vertex -3.919689 -4.242640 -1.623588
       vertex -2.582671 -5.000000 -1.623588
       vertex -2.872500 -5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816085
    outer loop
       vertex 4.242640 0.000000 4.242640
       vertex 3.109172 0.000000 5.000000
       vertex 2.414214 -1.979074 5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816085
    outer loop
       vertex 4.242640 0.000000 4.242640
       vertex 2.414214 -1.979074 5.000000
       vertex 2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 0.816085
    outer loop
       vertex 4.242640 0.000000 4.242640
       vertex 2.302890 -2.296101 5.000000
       vertex 3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal 0.552209 0.826439 -0.109841
    outer loop
       vertex 3.109172 5.000000 -0.000000
       vertex 4.242640 4.242640 -0.000000
       vertex 3.007397 5.000000 -0.511654
    endloop
  endfacet
  facet normal 0.552209 0.826439 -0.109841
    outer loop
       vertex 3.007397 5.000000 -0.511654
       vertex 4.242640 4.242640 -0.000000
       vertex 3.919689 4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.552209 0.826439 -0.109841
    outer loop
       vertex 3.007397 5.000000 -0.511654
       vertex 3.919689 4.242640 -1.623588
       vertex 2.872500 5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.464677 -0.548124 0.695439
    outer loop
       vertex 3.919689 -2.296101 3.919689
       vertex 2.302890 -2.296101 5.000000
       vertex 2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 0.695439
    outer loop
       vertex 3.919689 -2.296101 3.919689
       vertex 2.071068 -2.492630 5.000000
       vertex 1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.464677 -0.548124 0.695439
    outer loop
       vertex 3.919689 -2.296101 3.919689
       vertex 1.623588 -4.242640 3.919689
       vertex 3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.312801 0.826438 -0.468140
    outer loop
       vertex 2.198517 5.000000 -2.198517
       vertex 3.000000 4.242640 -3.000000
       vertex 1.764757 5.000000 -2.488345
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex 1.764757 5.000000 -2.488345
       vertex 3.000000 4.242640 -3.000000
       vertex 1.623588 4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.312801 0.826439 -0.468140
    outer loop
       vertex 1.764757 5.000000 -2.488345
       vertex 1.623588 4.242640 -3.919689
       vertex 1.189829 5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552208
    outer loop
       vertex -1.189829 5.000000 2.872500
       vertex -1.623588 4.242640 3.919689
       vertex -0.678174 5.000000 2.974275
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552209
    outer loop
       vertex -0.678174 5.000000 2.974275
       vertex -1.623588 4.242640 3.919689
       vertex 0.000000 4.242640 4.242640
    endloop
  endfacet
  facet normal -0.109841 0.826439 0.552209
    outer loop
       vertex -0.678174 5.000000 2.974275
       vertex 0.000000 4.242640 4.242640
       vertex 0.000000 5.000000 3.109172
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex 2.872500 5.000000 -1.189829
       vertex 3.919689 4.242640 -1.623588
       vertex 2.582671 5.000000 -1.623588
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex 2.582671 5.000000 -1.623588
       vertex 3.919689 4.242640 -1.623588
       vertex 3.000000 4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.468140 0.826439 -0.312801
    outer loop
       vertex 2.582671 5.000000 -1.623588
       vertex 3.000000 4.242640 -3.000000
       vertex 2.198517 5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex -3.000000 -4.242640 -3.000000
       vertex -1.623588 -4.242640 -3.919689
       vertex -1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal -0.312801 -0.826439 -0.468140
    outer loop
       vertex -3.000000 -4.242640 -3.000000
       vertex -1.189829 -5.000000 -2.872500
       vertex -1.764757 -5.000000 -2.488345
    endloop
  endfacet
  facet normal -0.312801 -0.826438 -0.468140
    outer loop
       vertex -3.000000 -4.242640 -3.000000
       vertex -1.764757 -5.000000 -2.488345
       vertex -2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex -1.623588 4.242640 3.919689
       vertex -3.000000 4.242640 3.000000
       vertex -3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex -1.623588 4.242640 3.919689
       vertex -3.919689 2.296101 3.919689
       vertex -2.302890 2.296101 5.000000
    endloop
  endfacet
  facet normal -0.464677 0.548124 0.695439
    outer loop
       vertex -1.623588 4.242640 3.919689
       vertex -2.302890 2.296101 5.000000
       vertex -2.071068 2.492630 5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex -2.302890 2.296101 5.000000
       vertex -3.919689 2.296101 3.919689
       vertex -4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex -2.302890 2.296101 5.000000
       vertex -4.242640 0.000000 4.242640
       vertex -2.414214 1.979074 5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 0.816085
    outer loop
       vertex -2.414214 1.979074 5.000000
       vertex -4.242640 0.000000 4.242640
       vertex -3.109172 0.000000 5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex -3.919689 2.296101 -3.919689
       vertex -2.302890 2.296101 -5.000000
       vertex -2.839378 0.768310 -5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816085
    outer loop
       vertex -3.919689 2.296101 -3.919689
       vertex -2.839378 0.768310 -5.000000
       vertex -3.109172 0.000000 -5.000000
    endloop
  endfacet
  facet normal -0.545290 0.191480 -0.816084
    outer loop
       vertex -3.919689 2.296101 -3.919689
       vertex -3.109172 0.000000 -5.000000
       vertex -4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex -0.000000 4.242640 -4.242640
       vertex 1.623588 4.242640 -3.919689
       vertex 2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex -0.000000 4.242640 -4.242640
       vertex 2.071068 2.492630 -5.000000
       vertex 1.828426 2.564862 -5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 -0.820326
    outer loop
       vertex -0.000000 4.242640 -4.242640
       vertex 1.828426 2.564862 -5.000000
       vertex -0.000000 3.109172 -5.000000
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex 3.919689 2.296101 -3.919689
       vertex 5.000000 2.296101 -2.302890
       vertex 5.000000 0.768310 -2.839378
    endloop
  endfacet
  facet normal 0.816085 0.191480 -0.545290
    outer loop
       vertex 3.919689 2.296101 -3.919689
       vertex 5.000000 0.768310 -2.839378
       vertex 5.000000 0.000000 -3.109172
    endloop
  endfacet
  facet normal 0.816084 0.191480 -0.545290
    outer loop
       vertex 3.919689 2.296101 -3.919689
       vertex 5.000000 0.000000 -3.109172
       vertex 4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal -0.552209 -0.826439 -0.109841
    outer loop
       vertex -4.242640 -4.242640 0.000000
       vertex -3.919689 -4.242640 -1.623588
       vertex -2.872500 -5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.552208 -0.826439 -0.109841
    outer loop
       vertex -4.242640 -4.242640 0.000000
       vertex -2.872500 -5.000000 -1.189829
       vertex -3.007397 -5.000000 -0.511654
    endloop
  endfacet
  facet normal -0.552209 -0.826439 -0.109841
    outer loop
       vertex -4.242640 -4.242640 0.000000
       vertex -3.007397 -5.000000 -0.511654
       vertex -3.109172 -5.000000 0.000000
    endloop
  endfacet
  facet normal -0.552208 0.826439 -0.109841
    outer loop
       vertex -2.872500 5.000000 -1.189829
       vertex -3.919689 4.242640 -1.623588
       vertex -2.974275 5.000000 -0.678174
    endloop
  endfacet
  facet normal -0.552209 0.826439 -0.109841
    outer loop
       vertex -2.974275 5.000000 -0.678174
       vertex -3.919689 4.242640 -1.623588
       vertex -4.242640 4.242640 0.000000
    endloop
  endfacet
  facet normal -0.552209 0.826439 -0.109841
    outer loop
       vertex -2.974275 5.000000 -0.678174
       vertex -4.242640 4.242640 0.000000
       vertex -3.109172 5.000000 0.000000
    endloop
  endfacet
  facet normal -0.816085 -0.191480 0.545290
    outer loop
       vertex -4.242640 0.000000 4.242640
       vertex -5.000000 0.000000 3.109172
       vertex -5.000000 -1.979074 2.414214
    endloop
  endfacet
  facet normal -0.816085 -0.191480 0.545290
    outer loop
       vertex -4.242640 0.000000 4.242640
       vertex -5.000000 -1.979074 2.414214
       vertex -5.000000 -2.296101 2.302890
    endloop
  endfacet
  facet normal -0.816085 -0.191480 0.545290
    outer loop
       vertex -4.242640 0.000000 4.242640
       vertex -5.000000 -2.296101 2.302890
       vertex -3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552209
    outer loop
       vertex 0.000000 -4.242640 4.242640
       vertex -1.623588 -4.242640 3.919689
       vertex -1.189829 -5.000000 2.872500
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552208
    outer loop
       vertex 0.000000 -4.242640 4.242640
       vertex -1.189829 -5.000000 2.872500
       vertex -0.511654 -5.000000 3.007397
    endloop
  endfacet
  facet normal -0.109841 -0.826439 0.552209
    outer loop
       vertex 0.000000 -4.242640 4.242640
       vertex -0.511654 -5.000000 3.007397
       vertex 0.000000 -5.000000 3.109172
    endloop
  endfacet
  facet normal -0.552209 -0.826439 0.109841
    outer loop
       vertex -3.919689 -4.242640 1.623588
       vertex -4.242640 -4.242640 0.000000
       vertex -3.109172 -5.000000 0.000000
    endloop
  endfacet
  facet normal -0.552209 -0.826439 0.109841
    outer loop
       vertex -3.919689 -4.242640 1.623588
       vertex -3.109172 -5.000000 0.000000
       vertex -2.974275 -5.000000 0.678174
    endloop
  endfacet
  facet normal -0.552208 -0.826439 0.109841
    outer loop
       vertex -3.919689 -4.242640 1.623588
       vertex -2.974275 -5.000000 0.678174
       vertex -2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex -2.302890 -2.296101 5.000000
       vertex -3.919689 -2.296101 3.919689
       vertex -3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex -2.302890 -2.296101 5.000000
       vertex -3.000000 -4.242640 3.000000
       vertex -2.171573 -2.407425 5.000000
    endloop
  endfacet
  facet normal -0.464677 -0.548124 0.695439
    outer loop
       vertex -2.171573 -2.407425 5.000000
       vertex -3.000000 -4.242640 3.000000
       vertex -1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal -0.464678 -0.548124 0.695439
    outer loop
       vertex -2.171573 -2.407425 5.000000
       vertex -1.623588 -4.242640 3.919689
       vertex -2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal -0.312801 0.826438 0.468140
    outer loop
       vertex -2.198517 5.000000 2.198517
       vertex -3.000000 4.242640 3.000000
       vertex -1.764757 5.000000 2.488345
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex -1.764757 5.000000 2.488345
       vertex -3.000000 4.242640 3.000000
       vertex -1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal -0.312801 0.826439 0.468140
    outer loop
       vertex -1.764757 5.000000 2.488345
       vertex -1.623588 4.242640 3.919689
       vertex -1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal -0.816084 -0.191480 -0.545290
    outer loop
       vertex -5.000000 0.000000 -3.109172
       vertex -4.242640 0.000000 -4.242640
       vertex -3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex -5.000000 0.000000 -3.109172
       vertex -3.919689 -2.296101 -3.919689
       vertex -5.000000 -0.768310 -2.839378
    endloop
  endfacet
  facet normal -0.816085 -0.191480 -0.545290
    outer loop
       vertex -5.000000 -0.768310 -2.839378
       vertex -3.919689 -2.296101 -3.919689
       vertex -5.000000 -2.296101 -2.302890
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816085
    outer loop
       vertex -4.242640 0.000000 -4.242640
       vertex -3.109172 0.000000 -5.000000
       vertex -2.414214 -1.979074 -5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816085
    outer loop
       vertex -4.242640 0.000000 -4.242640
       vertex -2.414214 -1.979074 -5.000000
       vertex -2.302890 -2.296101 -5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 -0.816085
    outer loop
       vertex -4.242640 0.000000 -4.242640
       vertex -2.302890 -2.296101 -5.000000
       vertex -3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex 3.000000 4.242640 -3.000000
       vertex 3.919689 4.242640 -1.623588
       vertex 5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex 3.000000 4.242640 -3.000000
       vertex 5.000000 2.492630 -2.071068
       vertex 5.000000 2.407425 -2.171573
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex 3.000000 4.242640 -3.000000
       vertex 5.000000 2.407425 -2.171573
       vertex 5.000000 2.296101 -2.302890
    endloop
  endfacet
  facet normal 0.695439 0.548124 -0.464677
    outer loop
       vertex 3.000000 4.242640 -3.000000
       vertex 5.000000 2.296101 -2.302890
       vertex 3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex 1.623588 4.242640 -3.919689
       vertex 3.000000 4.242640 -3.000000
       vertex 3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex 1.623588 4.242640 -3.919689
       vertex 3.919689 2.296101 -3.919689
       vertex 2.302890 2.296101 -5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 -0.695439
    outer loop
       vertex 1.623588 4.242640 -3.919689
       vertex 2.302890 2.296101 -5.000000
       vertex 2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex -5.000000 -3.109172 0.000000
       vertex -4.242640 -4.242640 0.000000
       vertex -5.000000 -2.564862 1.828426
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163173
    outer loop
       vertex -5.000000 -2.564862 1.828426
       vertex -4.242640 -4.242640 0.000000
       vertex -3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal -0.820326 -0.548124 0.163172
    outer loop
       vertex -5.000000 -2.564862 1.828426
       vertex -3.919689 -4.242640 1.623588
       vertex -5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex -3.919689 4.242640 1.623588
       vertex -4.242640 4.242640 0.000000
       vertex -5.000000 3.109172 0.000000
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex -3.919689 4.242640 1.623588
       vertex -5.000000 3.109172 0.000000
       vertex -5.000000 2.947442 0.543277
    endloop
  endfacet
  facet normal -0.820326 0.548124 0.163173
    outer loop
       vertex -3.919689 4.242640 1.623588
       vertex -5.000000 2.947442 0.543277
       vertex -5.000000 2.492630 2.071068
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex -2.071068 -2.492630 5.000000
       vertex -1.623588 -4.242640 3.919689
       vertex -0.543277 -2.947442 5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex -0.543277 -2.947442 5.000000
       vertex -1.623588 -4.242640 3.919689
       vertex 0.000000 -4.242640 4.242640
    endloop
  endfacet
  facet normal -0.163173 -0.548124 0.820326
    outer loop
       vertex -0.543277 -2.947442 5.000000
       vertex 0.000000 -4.242640 4.242640
       vertex 0.000000 -3.109172 5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816084
    outer loop
       vertex 3.109172 0.000000 -5.000000
       vertex 4.242640 0.000000 -4.242640
       vertex 3.919689 -2.296101 -3.919689
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex 3.109172 0.000000 -5.000000
       vertex 3.919689 -2.296101 -3.919689
       vertex 2.839378 -0.768310 -5.000000
    endloop
  endfacet
  facet normal 0.545290 -0.191480 -0.816085
    outer loop
       vertex 2.839378 -0.768310 -5.000000
       vertex 3.919689 -2.296101 -3.919689
       vertex 2.302890 -2.296101 -5.000000
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex -2.872500 5.000000 1.189829
       vertex -3.919689 4.242640 1.623588
       vertex -2.582671 5.000000 1.623588
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex -2.582671 5.000000 1.623588
       vertex -3.919689 4.242640 1.623588
       vertex -3.000000 4.242640 3.000000
    endloop
  endfacet
  facet normal -0.468140 0.826439 0.312801
    outer loop
       vertex -2.582671 5.000000 1.623588
       vertex -3.000000 4.242640 3.000000
       vertex -2.198517 5.000000 2.198517
    endloop
  endfacet
  facet normal -0.464677 -0.548124 -0.695439
    outer loop
       vertex -3.919689 -2.296101 -3.919689
       vertex -2.302890 -2.296101 -5.000000
       vertex -2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.464677 -0.548124 -0.695439
    outer loop
       vertex -3.919689 -2.296101 -3.919689
       vertex -2.071068 -2.492630 -5.000000
       vertex -1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.464677 -0.548124 -0.695439
    outer loop
       vertex -3.919689 -2.296101 -3.919689
       vertex -1.623588 -4.242640 -3.919689
       vertex -3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex 4.242640 4.242640 0.000000
       vertex 3.919689 4.242640 1.623588
       vertex 5.000000 2.492630 2.071068
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex 4.242640 4.242640 0.000000
       vertex 5.000000 2.492630 2.071068
       vertex 5.000000 2.564862 1.828426
    endloop
  endfacet
  facet normal 0.820326 0.548124 0.163173
    outer loop
       vertex 4.242640 4.242640 0.000000
       vertex 5.000000 2.564862 1.828426
       vertex 5.000000 3.109172 0.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 0.109841
    outer loop
       vertex 4.242640 -4.242640 0.000000
       vertex 3.919689 -4.242640 1.623588
       vertex 2.872500 -5.000000 1.189829
    endloop
  endfacet
  facet normal 0.552208 -0.826439 0.109841
    outer loop
       vertex 4.242640 -4.242640 0.000000
       vertex 2.872500 -5.000000 1.189829
       vertex 3.007397 -5.000000 0.511654
    endloop
  endfacet
  facet normal 0.552209 -0.826439 0.109841
    outer loop
       vertex 4.242640 -4.242640 0.000000
       vertex 3.007397 -5.000000 0.511654
       vertex 3.109172 -5.000000 0.000000
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552209
    outer loop
       vertex -0.000000 -4.242640 -4.242640
       vertex 1.623588 -4.242640 -3.919689
       vertex 1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552208
    outer loop
       vertex -0.000000 -4.242640 -4.242640
       vertex 1.189829 -5.000000 -2.872500
       vertex 0.511654 -5.000000 -3.007397
    endloop
  endfacet
  facet normal 0.109841 -0.826439 -0.552209
    outer loop
       vertex -0.000000 -4.242640 -4.242640
       vertex 0.511654 -5.000000 -3.007397
       vertex -0.000000 -5.000000 -3.109172
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex 0.000000 -3.109172 5.000000
       vertex 0.000000 -4.242640 4.242640
       vertex 1.828426 -2.564862 5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 0.820326
    outer loop
       vertex 1.828426 -2.564862 5.000000
       vertex 0.000000 -4.242640 4.242640
       vertex 1.623588 -4.242640 3.919689
    endloop
  endfacet
  facet normal 0.163172 -0.548124 0.820326
    outer loop
       vertex 1.828426 -2.564862 5.000000
       vertex 1.623588 -4.242640 3.919689
       vertex 2.071068 -2.492630 5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex 2.302890 2.296101 -5.000000
       vertex 3.919689 2.296101 -3.919689
       vertex 4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex 2.302890 2.296101 -5.000000
       vertex 4.242640 0.000000 -4.242640
       vertex 2.414214 1.979074 -5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 -0.816085
    outer loop
       vertex 2.414214 1.979074 -5.000000
       vertex 4.242640 0.000000 -4.242640
       vertex 3.109172 0.000000 -5.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 -0.109841
    outer loop
       vertex 3.919689 -4.242640 -1.623588
       vertex 4.242640 -4.242640 -0.000000
       vertex 3.109172 -5.000000 -0.000000
    endloop
  endfacet
  facet normal 0.552209 -0.826439 -0.109841
    outer loop
       vertex 3.919689 -4.242640 -1.623588
       vertex 3.109172 -5.000000 -0.000000
       vertex 2.974275 -5.000000 -0.678174
    endloop
  endfacet
  facet normal 0.552208 -0.826439 -0.109841
    outer loop
       vertex 3.919689 -4.242640 -1.623588
       vertex 2.974275 -5.000000 -0.678174
       vertex 2.872500 -5.000000 -1.189829
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex -1.623588 4.242640 -3.919689
       vertex -0.000000 4.242640 -4.242640
       vertex -0.000000 3.109172 -5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex -1.623588 4.242640 -3.919689
       vertex -0.000000 3.109172 -5.000000
       vertex -0.543277 2.947442 -5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 -0.820326
    outer loop
       vertex -1.623588 4.242640 -3.919689
       vertex -0.543277 2.947442 -5.000000
       vertex -2.071068 2.492630 -5.000000
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex 5.000000 -2.492630 2.071068
       vertex 3.919689 -4.242640 1.623588
       vertex 5.000000 -2.947442 0.543277
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex 5.000000 -2.947442 0.543277
       vertex 3.919689 -4.242640 1.623588
       vertex 4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal 0.820326 -0.548124 0.163173
    outer loop
       vertex 5.000000 -2.947442 0.543277
       vertex 4.242640 -4.242640 0.000000
       vertex 5.000000 -3.109172 0.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816084
    outer loop
       vertex -3.109172 0.000000 5.000000
       vertex -4.242640 0.000000 4.242640
       vertex -3.919689 -2.296101 3.919689
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex -3.109172 0.000000 5.000000
       vertex -3.919689 -2.296101 3.919689
       vertex -2.839378 -0.768310 5.000000
    endloop
  endfacet
  facet normal -0.545290 -0.191480 0.816085
    outer loop
       vertex -2.839378 -0.768310 5.000000
       vertex -3.919689 -2.296101 3.919689
       vertex -2.302890 -2.296101 5.000000
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex -5.000000 -2.296101 -2.302890
       vertex -3.919689 -2.296101 -3.919689
       vertex -3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex -5.000000 -2.296101 -2.302890
       vertex -3.000000 -4.242640 -3.000000
       vertex -5.000000 -2.407425 -2.171573
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464677
    outer loop
       vertex -5.000000 -2.407425 -2.171573
       vertex -3.000000 -4.242640 -3.000000
       vertex -3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal -0.695439 -0.548124 -0.464678
    outer loop
       vertex -5.000000 -2.407425 -2.171573
       vertex -3.919689 -4.242640 -1.623588
       vertex -5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex -4.242640 4.242640 0.000000
       vertex -3.919689 4.242640 -1.623588
       vertex -5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex -4.242640 4.242640 0.000000
       vertex -5.000000 2.492630 -2.071068
       vertex -5.000000 2.564862 -1.828426
    endloop
  endfacet
  facet normal -0.820326 0.548124 -0.163173
    outer loop
       vertex -4.242640 4.242640 0.000000
       vertex -5.000000 2.564862 -1.828426
       vertex -5.000000 3.109172 0.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex 2.302890 -2.296101 -5.000000
       vertex 3.919689 -2.296101 -3.919689
       vertex 3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex 2.302890 -2.296101 -5.000000
       vertex 3.000000 -4.242640 -3.000000
       vertex 2.171573 -2.407425 -5.000000
    endloop
  endfacet
  facet normal 0.464677 -0.548124 -0.695439
    outer loop
       vertex 2.171573 -2.407425 -5.000000
       vertex 3.000000 -4.242640 -3.000000
       vertex 1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal 0.464678 -0.548124 -0.695439
    outer loop
       vertex 2.171573 -2.407425 -5.000000
       vertex 1.623588 -4.242640 -3.919689
       vertex 2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex -5.000000 2.296101 -2.302890
       vertex -3.919689 2.296101 -3.919689
       vertex -4.242640 0.000000 -4.242640
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex -5.000000 2.296101 -2.302890
       vertex -4.242640 0.000000 -4.242640
       vertex -5.000000 1.979074 -2.414214
    endloop
  endfacet
  facet normal -0.816085 0.191480 -0.545290
    outer loop
       vertex -5.000000 1.979074 -2.414214
       vertex -4.242640 0.000000 -4.242640
       vertex -5.000000 0.000000 -3.109172
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464677
    outer loop
       vertex 3.919689 -2.296101 -3.919689
       vertex 5.000000 -2.296101 -2.302890
       vertex 5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464677
    outer loop
       vertex 3.919689 -2.296101 -3.919689
       vertex 5.000000 -2.492630 -2.071068
       vertex 3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.695439 -0.548124 -0.464677
    outer loop
       vertex 3.919689 -2.296101 -3.919689
       vertex 3.919689 -4.242640 -1.623588
       vertex 3.000000 -4.242640 -3.000000
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex 3.000000 -4.242640 3.000000
       vertex 1.623588 -4.242640 3.919689
       vertex 1.189829 -5.000000 2.872500
    endloop
  endfacet
  facet normal 0.312801 -0.826439 0.468140
    outer loop
       vertex 3.000000 -4.242640 3.000000
       vertex 1.189829 -5.000000 2.872500
       vertex 1.764757 -5.000000 2.488345
    endloop
  endfacet
  facet normal 0.312801 -0.826438 0.468140
    outer loop
       vertex 3.000000 -4.242640 3.000000
       vertex 1.764757 -5.000000 2.488345
       vertex 2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex 5.000000 -2.296101 2.302890
       vertex 3.919689 -2.296101 3.919689
       vertex 3.000000 -4.242640 3.000000
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex 5.000000 -2.296101 2.302890
       vertex 3.000000 -4.242640 3.000000
       vertex 5.000000 -2.407425 2.171573
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464677
    outer loop
       vertex 5.000000 -2.407425 2.171573
       vertex 3.000000 -4.242640 3.000000
       vertex 3.919689 -4.242640 1.623588
    endloop
  endfacet
  facet normal 0.695439 -0.548124 0.464678
    outer loop
       vertex 5.000000 -2.407425 2.171573
       vertex 3.919689 -4.242640 1.623588
       vertex 5.000000 -2.492630 2.071068
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex 3.919689 2.296101 3.919689
       vertex 2.302890 2.296101 5.000000
       vertex 2.839378 0.768310 5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816085
    outer loop
       vertex 3.919689 2.296101 3.919689
       vertex 2.839378 0.768310 5.000000
       vertex 3.109172 0.000000 5.000000
    endloop
  endfacet
  facet normal 0.545290 0.191480 0.816084
    outer loop
       vertex 3.919689 2.296101 3.919689
       vertex 3.109172 0.000000 5.000000
       vertex 4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex -3.919689 4.242640 -1.623588
       vertex -3.000000 4.242640 -3.000000
       vertex -3.919689 2.296101 -3.919689
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex -3.919689 4.242640 -1.623588
       vertex -3.919689 2.296101 -3.919689
       vertex -5.000000 2.296101 -2.302890
    endloop
  endfacet
  facet normal -0.695439 0.548124 -0.464677
    outer loop
       vertex -3.919689 4.242640 -1.623588
       vertex -5.000000 2.296101 -2.302890
       vertex -5.000000 2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex 0.000000 4.242640 4.242640
       vertex -1.623588 4.242640 3.919689
       vertex -2.071068 2.492630 5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex 0.000000 4.242640 4.242640
       vertex -2.071068 2.492630 5.000000
       vertex -1.828426 2.564862 5.000000
    endloop
  endfacet
  facet normal -0.163173 0.548124 0.820326
    outer loop
       vertex 0.000000 4.242640 4.242640
       vertex -1.828426 2.564862 5.000000
       vertex 0.000000 3.109172 5.000000
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex -3.000000 4.242640 3.000000
       vertex -3.919689 4.242640 1.623588
       vertex -5.000000 2.492630 2.071068
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex -3.000000 4.242640 3.000000
       vertex -5.000000 2.492630 2.071068
       vertex -5.000000 2.407425 2.171573
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex -3.000000 4.242640 3.000000
       vertex -5.000000 2.407425 2.171573
       vertex -5.000000 2.296101 2.302890
    endloop
  endfacet
  facet normal -0.695439 0.548124 0.464677
    outer loop
       vertex -3.000000 4.242640 3.000000
       vertex -5.000000 2.296101 2.302890
       vertex -3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex -0.000000 -3.109172 -5.000000
       vertex -0.000000 -4.242640 -4.242640
       vertex -1.828426 -2.564862 -5.000000
    endloop
  endfacet
  facet normal -0.163173 -0.548124 -0.820326
    outer loop
       vertex -1.828426 -2.564862 -5.000000
       vertex -0.000000 -4.242640 -4.242640
       vertex -1.623588 -4.242640 -3.919689
    endloop
  endfacet
  facet normal -0.163172 -0.548124 -0.820326
    outer loop
       vertex -1.828426 -2.564862 -5.000000
       vertex -1.623588 -4.242640 -3.919689
       vertex -2.071068 -2.492630 -5.000000
    endloop
  endfacet
  facet normal -0.552209 0.826439 0.109841
    outer loop
       vertex -3.109172 5.000000 0.000000
       vertex -4.242640 4.242640 0.000000
       vertex -3.007397 5.000000 0.511654
    endloop
  endfacet
  facet normal -0.552209 0.826439 0.109841
    outer loop
       vertex -3.007397 5.000000 0.511654
       vertex -4.242640 4.242640 0.000000
       vertex -3.919689 4.242640 1.623588
    endloop
  endfacet
  facet normal -0.552209 0.826439 0.109841
    outer loop
       vertex -3.007397 5.000000 0.511654
       vertex -3.919689 4.242640 1.623588
       vertex -2.872500 5.000000 1.189829
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex 3.000000 4.242640 3.000000
       vertex 1.623588 4.242640 3.919689
       vertex 2.071068 2.492630 5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex 3.000000 4.242640 3.000000
       vertex 2.071068 2.492630 5.000000
       vertex 2.171573 2.407425 5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex 3.000000 4.242640 3.000000
       vertex 2.171573 2.407425 5.000000
       vertex 2.302890 2.296101 5.000000
    endloop
  endfacet
  facet normal 0.464677 0.548124 0.695439
    outer loop
       vertex 3.000000 4.242640 3.000000
       vertex 2.302890 2.296101 5.000000
       vertex 3.919689 2.296101 3.919689
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.000000 -2.492630 -2.071068
       vertex -3.919689 -4.242640 -1.623588
       vertex -5.000000 -2.947442 -0.543277
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.000000 -2.947442 -0.543277
       vertex -3.919689 -4.242640 -1.623588
       vertex -4.242640 -4.242640 0.000000
    endloop
  endfacet
  facet normal -0.820326 -0.548124 -0.163173
    outer loop
       vertex -5.000000 -2.947442 -0.543277
       vertex -4.242640 -4.242640 0.000000
       vertex -5.000000 -3.109172 0.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex 1.623588 4.242640 3.919689
       vertex 0.000000 4.242640 4.242640
       vertex 0.000000 3.109172 5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex 1.623588 4.242640 3.919689
       vertex 0.000000 3.109172 5.000000
       vertex 0.543277 2.947442 5.000000
    endloop
  endfacet
  facet normal 0.163173 0.548124 0.820326
    outer loop
       vertex 1.623588 4.242640 3.919689
       vertex 0.543277 2.947442 5.000000
       vertex 2.071068 2.492630 5.000000
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex 1.623588 -4.242640 -3.919689
       vertex 3.000000 -4.242640 -3.000000
       vertex 2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex 1.623588 -4.242640 -3.919689
       vertex 2.198517 -5.000000 -2.198517
       vertex 1.623588 -5.000000 -2.582671
    endloop
  endfacet
  facet normal 0.312801 -0.826439 -0.468140
    outer loop
       vertex 1.623588 -4.242640 -3.919689
       vertex 1.623588 -5.000000 -2.582671
       vertex 1.189829 -5.000000 -2.872500
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex 2.071068 -2.492630 -5.000000
       vertex 1.623588 -4.242640 -3.919689
       vertex 0.543277 -2.947442 -5.000000
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex 0.543277 -2.947442 -5.000000
       vertex 1.623588 -4.242640 -3.919689
       vertex -0.000000 -4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.163173 -0.548124 -0.820326
    outer loop
       vertex 0.543277 -2.947442 -5.000000
       vertex -0.000000 -4.242640 -4.242640
       vertex -0.000000 -3.109172 -5.000000
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552209
    outer loop
       vertex 0.000000 5.000000 3.109172
       vertex 0.000000 4.242640 4.242640
       vertex 0.511654 5.000000 3.007397
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552209
    outer loop
       vertex 0.511654 5.000000 3.007397
       vertex 0.000000 4.242640 4.242640
       vertex 1.623588 4.242640 3.919689
    endloop
  endfacet
  facet normal 0.109841 0.826439 0.552209
    outer loop
       vertex 0.511654 5.000000 3.007397
       vertex 1.623588 4.242640 3.919689
       vertex 1.189829 5.000000 2.872500
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex 3.000000 -4.242640 -3.000000
       vertex 3.919689 -4.242640 -1.623588
       vertex 2.872500 -5.000000 -1.189829
    endloop
  endfacet
  facet normal 0.468140 -0.826439 -0.312801
    outer loop
       vertex 3.000000 -4.242640 -3.000000
       vertex 2.872500 -5.000000 -1.189829
       vertex 2.488345 -5.000000 -1.764757
    endloop
  endfacet
  facet normal 0.468140 -0.826438 -0.312801
    outer loop
       vertex 3.000000 -4.242640 -3.000000
       vertex 2.488345 -5.000000 -1.764757
       vertex 2.198517 -5.000000 -2.198517
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex -3.919689 2.296101 3.919689
       vertex -5.000000 2.296101 2.302890
       vertex -5.000000 0.768310 2.839378
    endloop
  endfacet
  facet normal -0.816085 0.191480 0.545290
    outer loop
       vertex -3.919689 2.296101 3.919689
       vertex -5.000000 0.768310 2.839378
       vertex -5.000000 0.000000 3.109172
    endloop
  endfacet
  facet normal -0.816084 0.191480 0.545290
    outer loop
       vertex -3.919689 2.296101 3.919689
       vertex -5.000000 0.000000 3.109172
       vertex -4.242640 0.000000 4.242640
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552209
    outer loop
       vertex 1.623588 -4.242640 3.919689
       vertex 0.000000 -4.242640 4.242640
       vertex 0.000000 -5.000000 3.109172
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552209
    outer loop
       vertex 1.623588 -4.242640 3.919689
       vertex 0.000000 -5.000000 3.109172
       vertex 0.678174 -5.000000 2.974275
    endloop
  endfacet
  facet normal 0.109841 -0.826439 0.552208
    outer loop
       vertex 1.623588 -4.242640 3.919689
       vertex 0.678174 -5.000000 2.974275
       vertex 1.189829 -5.000000 2.872500
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552208
    outer loop
       vertex 1.189829 5.000000 -2.872500
       vertex 1.623588 4.242640 -3.919689
       vertex 0.678174 5.000000 -2.974275
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552209
    outer loop
       vertex 0.678174 5.000000 -2.974275
       vertex 1.623588 4.242640 -3.919689
       vertex -0.000000 4.242640 -4.242640
    endloop
  endfacet
  facet normal 0.109841 0.826439 -0.552209
    outer loop
       vertex 0.678174 5.000000 -2.974275
       vertex -0.000000 4.242640 -4.242640
       vertex -0.000000 5.000000 -3.109172
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex 5.000000 -3.109172 -0.000000
       vertex 4.242640 -4.242640 -0.000000
       vertex 5.000000 -2.564862 -1.828426
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163173
    outer loop
       vertex 5.000000 -2.564862 -1.828426
       vertex 4.242640 -4.242640 -0.000000
       vertex 3.919689 -4.242640 -1.623588
    endloop
  endfacet
  facet normal 0.820326 -0.548124 -0.163172
    outer loop
       vertex 5.000000 -2.564862 -1.828426
       vertex 3.919689 -4.242640 -1.623588
       vertex 5.000000 -2.492630 -2.071068
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex -1.623588 -4.242640 3.919689
       vertex -3.000000 -4.242640 3.000000
       vertex -2.198517 -5.000000 2.198517
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex -1.623588 -4.242640 3.919689
       vertex -2.198517 -5.000000 2.198517
       vertex -1.623588 -5.000000 2.582671
    endloop
  endfacet
  facet normal -0.312801 -0.826439 0.468140
    outer loop
       vertex -1.623588 -4.242640 3.919689
       vertex -1.623588 -5.000000 2.582671
       vertex -1.189829 -5.000000 2.872500
    endloop
  endfacet
endsolid
```

### Hull

```stl
solid ScadSharp
  facet normal -0.001067 0.195090 0.980785
    outer loop
       vertex 5.000000 -5.000000 -5.000000
       vertex -0.000000 0.000000 -6.000000
       vertex -0.000000 -2.296101 -5.543277
    endloop
  endfacet
  facet normal 0.195090 -0.001067 0.980785
    outer loop
       vertex -5.000000 5.000000 -5.000000
       vertex -2.296101 0.000000 -5.543277
       vertex -0.000000 0.000000 -6.000000
    endloop
  endfacet
  facet normal -0.196987 -0.980406 -0.000000
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 2.296101 5.543277 0.000000
       vertex 5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.001067 0.195090 -0.980785
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex 0.000000 0.000000 6.000000
       vertex 0.000000 -2.296101 5.543277
    endloop
  endfacet
  facet normal 0.000000 0.980406 -0.196987
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 0.000000 -5.543277 2.296101
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal -0.001067 -0.195090 0.980785
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex -0.000000 2.296101 -5.543277
       vertex -0.000000 0.000000 -6.000000
    endloop
  endfacet
  facet normal 0.001067 0.195090 0.980785
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -0.000000 -2.296101 -5.543277
       vertex -0.000000 0.000000 -6.000000
    endloop
  endfacet
  facet normal -0.980785 0.001067 0.195090
    outer loop
       vertex 5.000000 -5.000000 -5.000000
       vertex 6.000000 0.000000 0.000000
       vertex 5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal -0.980785 -0.001067 -0.195090
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 6.000000 0.000000 0.000000
       vertex 5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal -0.195090 -0.001067 -0.980785
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 2.296101 0.000000 5.543277
       vertex 0.000000 0.000000 6.000000
    endloop
  endfacet
  facet normal -0.980785 0.195090 0.001067
    outer loop
       vertex 5.000000 -5.000000 -5.000000
       vertex 5.543277 -2.296101 0.000000
       vertex 6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal 0.195090 0.001067 0.980785
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -0.000000 0.000000 -6.000000
       vertex -2.296101 0.000000 -5.543277
    endloop
  endfacet
  facet normal -0.980406 0.000000 0.196987
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 5.000000 -5.000000 -5.000000
       vertex 5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal -0.980785 0.001067 -0.195090
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 5.543277 0.000000 2.296101
       vertex 6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal -0.980785 0.195090 -0.001067
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 6.000000 0.000000 0.000000
       vertex 5.543277 -2.296101 0.000000
    endloop
  endfacet
  facet normal 0.196987 -0.000000 0.980406
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -2.296101 0.000000 -5.543277
       vertex -5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.001067 -0.195090 0.980785
    outer loop
       vertex -5.000000 5.000000 -5.000000
       vertex -0.000000 0.000000 -6.000000
       vertex -0.000000 2.296101 -5.543277
    endloop
  endfacet
  facet normal 0.980406 0.000000 -0.196987
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 -5.000000 5.000000
       vertex -5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal 0.980785 -0.001067 0.195090
    outer loop
       vertex -5.000000 5.000000 -5.000000
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal -0.980406 0.000000 -0.196987
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 5.000000 5.000000
       vertex 5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal -0.196987 0.000000 0.980406
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 2.296101 0.000000 -5.543277
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.196987 0.000000 -0.980406
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -2.296101 0.000000 5.543277
       vertex -5.000000 -5.000000 5.000000
    endloop
  endfacet
  facet normal 0.980406 0.196987 0.000000
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.543277 -2.296101 0.000000
    endloop
  endfacet
  facet normal -0.196987 0.980406 0.000000
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 5.000000 -5.000000 -5.000000
       vertex 2.296101 -5.543277 0.000000
    endloop
  endfacet
  facet normal -0.980406 0.196987 0.000000
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 5.543277 -2.296101 0.000000
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.195090 0.980785 -0.001067
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex 0.000000 -6.000000 0.000000
       vertex -2.296101 -5.543277 0.000000
    endloop
  endfacet
  facet normal 0.980406 -0.000000 0.196987
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.000000 5.000000 -5.000000
       vertex -5.543277 0.000000 -2.296101
    endloop
  endfacet
  facet normal 0.001067 0.980785 0.195090
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex 0.000000 -6.000000 0.000000
       vertex -0.000000 -5.543277 -2.296101
    endloop
  endfacet
  facet normal -0.980785 -0.001067 0.195090
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 5.543277 0.000000 -2.296101
       vertex 6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal 0.195090 -0.980785 -0.001067
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -2.296101 5.543277 0.000000
       vertex 0.000000 6.000000 0.000000
    endloop
  endfacet
  facet normal 0.980785 0.195090 0.001067
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 -2.296101 0.000000
    endloop
  endfacet
  facet normal 0.001067 -0.195090 -0.980785
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex 0.000000 2.296101 5.543277
       vertex 0.000000 0.000000 6.000000
    endloop
  endfacet
  facet normal -0.196987 0.000000 -0.980406
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 2.296101 0.000000 5.543277
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.195090 0.001067 -0.980785
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex -2.296101 0.000000 5.543277
       vertex 0.000000 0.000000 6.000000
    endloop
  endfacet
  facet normal 0.980785 -0.001067 -0.195090
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -5.543277 0.000000 2.296101
       vertex -6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal 0.000000 -0.980406 0.196987
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex -0.000000 5.543277 -2.296101
       vertex -5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal -0.000000 0.980406 0.196987
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -0.000000 -5.543277 -2.296101
       vertex 5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal -0.195090 0.001067 0.980785
    outer loop
       vertex 5.000000 -5.000000 -5.000000
       vertex 2.296101 0.000000 -5.543277
       vertex -0.000000 0.000000 -6.000000
    endloop
  endfacet
  facet normal 0.000000 -0.980406 -0.196987
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex 0.000000 5.543277 2.296101
       vertex 5.000000 5.000000 5.000000
    endloop
  endfacet
  facet normal 0.980785 0.195090 -0.001067
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex -5.543277 -2.296101 0.000000
       vertex -6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal 0.196987 -0.980406 0.000000
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -5.000000 5.000000 -5.000000
       vertex -2.296101 5.543277 0.000000
    endloop
  endfacet
  facet normal -0.980785 -0.195090 0.001067
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 6.000000 0.000000 0.000000
       vertex 5.543277 2.296101 0.000000
    endloop
  endfacet
  facet normal 0.195090 0.980785 0.001067
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -2.296101 -5.543277 0.000000
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal -0.980785 -0.195090 -0.001067
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 5.543277 2.296101 0.000000
       vertex 6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal -0.001067 0.980785 -0.195090
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 0.000000 -6.000000 0.000000
       vertex 0.000000 -5.543277 2.296101
    endloop
  endfacet
  facet normal 0.196987 0.980406 0.000000
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex -2.296101 -5.543277 0.000000
       vertex -5.000000 -5.000000 -5.000000
    endloop
  endfacet
  facet normal -0.195090 -0.001067 0.980785
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex -0.000000 0.000000 -6.000000
       vertex 2.296101 0.000000 -5.543277
    endloop
  endfacet
  facet normal -0.001067 -0.980785 -0.195090
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 0.000000 5.543277 2.296101
       vertex 0.000000 6.000000 0.000000
    endloop
  endfacet
  facet normal 0.980785 -0.195090 -0.001067
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 2.296101 0.000000
    endloop
  endfacet
  facet normal -0.195090 0.980785 0.001067
    outer loop
       vertex 5.000000 -5.000000 -5.000000
       vertex 0.000000 -6.000000 0.000000
       vertex 2.296101 -5.543277 0.000000
    endloop
  endfacet
  facet normal 0.195090 -0.001067 -0.980785
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex 0.000000 0.000000 6.000000
       vertex -2.296101 0.000000 5.543277
    endloop
  endfacet
  facet normal 0.001067 -0.980785 0.195090
    outer loop
       vertex -5.000000 5.000000 -5.000000
       vertex -0.000000 5.543277 -2.296101
       vertex 0.000000 6.000000 0.000000
    endloop
  endfacet
  facet normal 0.000000 0.196987 -0.980406
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex -5.000000 -5.000000 5.000000
       vertex 0.000000 -2.296101 5.543277
    endloop
  endfacet
  facet normal -0.195090 -0.980785 0.001067
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 2.296101 5.543277 0.000000
       vertex 0.000000 6.000000 0.000000
    endloop
  endfacet
  facet normal -0.195090 -0.980785 -0.001067
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 0.000000 6.000000 0.000000
       vertex 2.296101 5.543277 0.000000
    endloop
  endfacet
  facet normal 0.001067 -0.980785 -0.195090
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex 0.000000 6.000000 0.000000
       vertex 0.000000 5.543277 2.296101
    endloop
  endfacet
  facet normal 0.001067 0.980785 -0.195090
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex 0.000000 -5.543277 2.296101
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.980785 0.001067 -0.195090
    outer loop
       vertex -5.000000 -5.000000 5.000000
       vertex -6.000000 0.000000 0.000000
       vertex -5.543277 0.000000 2.296101
    endloop
  endfacet
  facet normal 0.000000 -0.196987 0.980406
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex -5.000000 5.000000 -5.000000
       vertex -0.000000 2.296101 -5.543277
    endloop
  endfacet
  facet normal 0.195090 -0.980785 0.001067
    outer loop
       vertex -5.000000 5.000000 -5.000000
       vertex 0.000000 6.000000 0.000000
       vertex -2.296101 5.543277 0.000000
    endloop
  endfacet
  facet normal -0.001067 0.195090 -0.980785
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 0.000000 -2.296101 5.543277
       vertex 0.000000 0.000000 6.000000
    endloop
  endfacet
  facet normal -0.195090 0.980785 -0.001067
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 2.296101 -5.543277 0.000000
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.000000 -0.196987 -0.980406
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex 5.000000 5.000000 5.000000
       vertex 0.000000 2.296101 5.543277
    endloop
  endfacet
  facet normal -0.195090 0.001067 -0.980785
    outer loop
       vertex 5.000000 -5.000000 5.000000
       vertex 0.000000 0.000000 6.000000
       vertex 2.296101 0.000000 5.543277
    endloop
  endfacet
  facet normal -0.001067 -0.195090 -0.980785
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 0.000000 0.000000 6.000000
       vertex 0.000000 2.296101 5.543277
    endloop
  endfacet
  facet normal 0.980406 -0.196987 0.000000
    outer loop
       vertex -5.000000 5.000000 5.000000
       vertex -5.543277 2.296101 0.000000
       vertex -5.000000 5.000000 -5.000000
    endloop
  endfacet
  facet normal 0.980785 -0.195090 0.001067
    outer loop
       vertex -5.000000 5.000000 -5.000000
       vertex -5.543277 2.296101 0.000000
       vertex -6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal -0.980406 -0.196987 -0.000000
    outer loop
       vertex 5.000000 5.000000 5.000000
       vertex 5.000000 5.000000 -5.000000
       vertex 5.543277 2.296101 0.000000
    endloop
  endfacet
  facet normal -0.001067 -0.980785 0.195090
    outer loop
       vertex 5.000000 5.000000 -5.000000
       vertex 0.000000 6.000000 0.000000
       vertex -0.000000 5.543277 -2.296101
    endloop
  endfacet
  facet normal -0.001067 0.980785 0.195090
    outer loop
       vertex 5.000000 -5.000000 -5.000000
       vertex -0.000000 -5.543277 -2.296101
       vertex 0.000000 -6.000000 0.000000
    endloop
  endfacet
  facet normal 0.980785 0.001067 0.195090
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex -5.543277 0.000000 -2.296101
       vertex -6.000000 0.000000 0.000000
    endloop
  endfacet
  facet normal -0.000000 0.196987 0.980406
    outer loop
       vertex -5.000000 -5.000000 -5.000000
       vertex 5.000000 -5.000000 -5.000000
       vertex -0.000000 -2.296101 -5.543277
    endloop
  endfacet
endsolid
```
