#include <complex.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct fileRead {
    char *fbuffer;
    long file_size;
};

struct fileRead read_file() {
    FILE *file;
    char *fbuffer;
    long file_size;

    file = fopen("./c/2016/inputs/day01a.txt", "rb");
    if (file == NULL) {
        perror("Error opening file");
    }

    fseek(file, 0, SEEK_END);
    file_size = ftell(file);
    rewind(file);

    fbuffer = (char *)malloc(sizeof(char) * file_size + 1);
    if (fbuffer == NULL) {
        perror("Memory error");
        fclose(file);
    }

    fread(fbuffer, 1, file_size, file);

    struct fileRead file_data = {fbuffer, file_size};

    fclose(file);

    return file_data;
}

int main() {
    struct fileRead file_data = read_file();
    long file_size = file_data.file_size;
    char *fbuffer = file_data.fbuffer;

    char *temp_buffer = (char *)malloc(sizeof(char) * (1));
    long temp_buffer_size = 0;
    int coords[2];
    coords[0] = 0;
    coords[1] = 0;
    int current_direction = 0;
    int record_cords[1024][2];
    int records_counter = 0;

    int found_visited_twice = 0;

    for (int i = 0; i < file_size; i++) {
        if (fbuffer[i] == ',') {
            char *temp_buffer =
                (char *)malloc(sizeof(char) * (temp_buffer_size + 1));
            strncpy(temp_buffer, fbuffer + i - temp_buffer_size,
                    temp_buffer_size);
            temp_buffer[sizeof temp_buffer] = '\0';
            char direction;
            int steps;

            if (sscanf(temp_buffer, "%c%d", &direction, &steps) != 2) {
                fprintf(stderr, "Error parsing temp_buffer\n");
                return EXIT_FAILURE;
            }

            if (direction == 'R') {
                current_direction++;
            } else {
                current_direction--;
            }

            if (current_direction > 3) {
                current_direction = 0;
            }

            if (current_direction < 0) {
                current_direction = 3;
            }

            for (int j = 0; j < steps; j++) {
                switch (current_direction) {
                    case 0:  // north
                        coords[0]++;
                        break;

                    case 1:  // east
                        coords[1]++;
                        break;

                    case 2:  // south
                        coords[0]--;
                        break;

                    case 3:  // west
                        coords[1]--;
                        break;
                }

                record_cords[records_counter][0] = coords[0];
                record_cords[records_counter][1] = coords[1];
				records_counter++;
            }

            i++;
            // records_counter++;
            temp_buffer_size = 0;
        } else {
            temp_buffer_size++;
        }
    }

	int position_found_twice = 0;
    for (int i = 0; i < records_counter; i++) {
        for (int j = 0; j < records_counter; j++) {
            if (i != j && record_cords[i][0] == record_cords[j][0] &&
                record_cords[i][1] == record_cords[j][1]) {
                printf("REPEATED LOCATION: %d %d\n", record_cords[i][0],
                       record_cords[i][1]);
				position_found_twice = i;
				break;
            }
			if (position_found_twice != 0) {
				break;
			}
        }
    }

    printf("PART ONE ANSWER: %d\n", abs(coords[0]) + abs(coords[1]));
    printf("PART TWO ANSWER: %d\n",
           abs(record_cords[position_found_twice][0] + abs(record_cords[position_found_twice][1])));

    free(file_data.fbuffer);

    return 0;
}
