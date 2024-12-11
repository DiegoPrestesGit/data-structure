#include <stdio.h>
#include <stdlib.h>
struct file_read {
    char *fb;
    long file_size;
};

struct file_read read_file(char *path) {
    FILE *file;
    char *fb;
    long file_size;

    file = fopen(path, "rb");
    if (file == NULL) {
        perror("Error opening file");
    }

    fseek(file, 0, SEEK_END);
    file_size = ftell(file);
    rewind(file);

    fb = (char *)malloc(sizeof(char) * file_size + 1);
    if (fb == NULL) {
        perror("Memory error");
        fclose(file);
    }

    fread(fb, 1, file_size, file);
    struct file_read file_data = {fb, file_size};
    fclose(file);
    return file_data;
}
