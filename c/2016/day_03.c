#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "utils.c"

// an array as a func param is a pointer, always inform the size of the array in
// another param
void bubbleSort(int *a, size_t arr_size) {
    for (int i = 0; i < arr_size - 1; i++) {
        for (int j = i + 1; j < arr_size; j++) {
            if (a[i] > a[j]) {
                int t = a[i];
                a[i] = a[j];
                a[j] = t;
            }
        }
    }
}

void a() {
    struct file_read f = read_file("./c/2016/inputs/day03.txt");

    int triangle_counter = 0;

    int a[3] = {0, 0, 0};
    char *n = malloc(8);
    int a_count = 0;
    for (int i = 0; i < f.file_size; i++) {
        switch (f.fb[i]) {
            case ' ': {
                if (strcmp(n, "")) {
                    break;
                }
                // xxXXXx
                // __123_321_99
                printf("%s", n);
                int num = atoi(n);
                a[a_count] = num;
                a_count++;
                n = "";

                break;
            }
            case '\n': {
                printf("%d", a[0]);
                printf("%s", " ");
                printf("%d", a[1]);
                printf("%s", " ");
                printf("%d", a[2]);
                printf("%s", "\n");

                if (a[0] == 0) {
                    break;
                }
                int num = atoi(n);
                a[a_count] = num;
                a_count++;
                n = "";

                if (a[1] != 0 && a[2] != 0) {
                    size_t s = sizeof(a) / sizeof(a[0]);
                    bubbleSort(a, s);
                    if (a[0] + a[1] > a[2]) {
                        triangle_counter++;
                    }
                }

                a[0] = 0;
                a[1] = 0;
                a[2] = 0;
                a_count = 0;
                break;
            }
            default: {
                char t[3];
                sprintf(t, "%d", f.fb[i]);
                strcat(n, t);
                printf("%s", n);
                break;
            }
        }
    }

    printf("PART 1: ");
    printf("%d", triangle_counter);
    printf("%c", '\n');

    free(f.fb);
}

int main() {
    a();
    return 0;
}
