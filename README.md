# Bilinear Gradient Generator

Uses bilinear-interpolation to generate random colour gradients. A random value is selected for each corner of the image, then bilinear-interpolation is used to interpolate the remaining pixels. This is repeated 3 times. These 3 separate pixel arrays are then combined into the red, green and blue channels to produce a final image.

## Examples 

![Example 1](examples/image1.png)

![Example 2](examples/image2.png)

![Example 3](examples/image3.png)

![Example 4](examples/image4.png)

![Example 5](examples/image5.png)

![Example 6](examples/image6.png)
