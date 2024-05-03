# GIF Image Resizer

This is a simple Go application that allows you to select a folder using a file dialog and choose values using sliders to resize GIF images.

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone git@github.com:simonof/gif-resizer.git
    ```

2. Navigate to the project directory:

    ```bash
    cd your-repo
    ```

3. Build and run the application:

    ```bash
    go run main.go
    ```

## Usage

1. When you run the application, a window will open with a button labeled "Open Folder".

2. Click the "Open Folder" button to select a folder containing GIF images using the file dialog that appears.

3. Three sliders will also be displayed. Use these sliders to select integer values between 1 and 999. These values represent the desired width, height, and frame count for resizing the GIF images.

4. Close the window after making your selections.

5. The application will process the GIF images in the selected folder, resizing them according to the selected values. The resized images will be saved in an output directory.

## Dependencies

- [Fyne](https://github.com/fyne-io/fyne): A cross-platform GUI toolkit in Go.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
