#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include "stdio.h"

struct file_read {
    char *fb;
    long file_size;
};

struct file_read read_file() {
    FILE *file;
    char *fb;
    long file_size;

    file = fopen("./c/2016/inputs/day02a.txt", "rb");
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

void a() {
    int m[3][3] = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
    struct file_read f = read_file();

    char *passwd = malloc(1);
    passwd[0] = '\0';

    int curr[2] = {1, 1};
    for (int i = 0; i < f.file_size; i++) {
        if (f.fb[i] == '\n') {
            char t[32];
            sprintf(t, "%d", m[curr[0]][curr[1]]);
            size_t new_len = strlen(passwd) + strlen(t) + 1;
            passwd = realloc(passwd, new_len);
            strcat(passwd, t);
            curr[0] = 1;
            curr[1] = 1;
            continue;
        }

        switch (f.fb[i]) {
            case 'U': {
                if (curr[0] != 0) {
                    curr[0] -= 1;
                }
                break;
            }

            case 'L': {
                if (curr[1] != 0) {
                    curr[1] -= 1;
                }
                break;
            }

            case 'R': {
                if (curr[1] != 2) {
                    curr[1] += 1;
                }
                break;
            }

            case 'D': {
                if (curr[0] != 2) {
                    curr[0] += 1;
                }
                break;
            }
        }
    }
    printf("PART 1: ");
    puts(passwd);
    free(passwd);
}

void b() {
    struct file_read f = read_file();
    char m[5][5] = {{'0', '0', '1', '0', '0'},
                    {'0', '2', '3', '4', '0'},
                    {'5', '6', '7', '8', '9'},
                    {'0', 'A', 'B', 'C', '0'},
                    {'0', '0', 'D', '0', '0'}};

    char *passwd = malloc(5);
    // passwd[0] = '\0';
    int count = 0;

    int curr[2] = {2, 0};
    for (int i = 0; i < f.file_size; i++) {
        if (f.fb[i] == '\n') {
            char t = m[curr[0]][curr[1]];
            passwd[count] = t;
            count++;
            continue;
        }

        switch (f.fb[i]) {
            case 'U': {
                if (curr[0] != 0 && m[curr[0] - 1][curr[1]] != '0') {
                    curr[0] -= 1;
                }
                break;
            }

            case 'L': {
                if (curr[1] != 0 && m[curr[0]][curr[1] - 1] != '0') {
                    curr[1] -= 1;
                }
                break;
            }

            case 'R': {
                if (curr[1] != 4 && m[curr[0]][curr[1] + 1] != '0') {
                    curr[1] += 1;
                }
                break;
            }

            case 'D': {
                if (curr[0] != 4 && m[curr[0] + 1][curr[1]] != '0') {
                    curr[0] += 1;
                }
                break;
            }
        }
    }
    printf("PART 2: ");
    printf("%s", passwd);
    printf("%s", "\n");
    free(passwd);
}

int main() {
    a();
    b();

    return 0;
}
