package day03;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;

public class A {
	public static void main(String[] args) throws IOException {
		String filePath = "./input.txt";
		String stringFromFile = stringFromFile(filePath);

		ArrayList<String> santaPath = new ArrayList<String>(Arrays.asList(stringFromFile.split("")));
		// POSITION: latitude | longitude
		ArrayList<int[]> positions = new ArrayList<int[]>(Arrays.asList(new int[] { 0, 0 }));
		int[] currPosition = { 0, 0 };

		for (int i = 0; i < santaPath.size(); i++) {
			switch (santaPath.get(i)) {
				case "^":
					currPosition = setCurrentPosition(currPosition, 1, 0);
					break;
				case "v":
					currPosition = setCurrentPosition(currPosition, -1, 0);
					break;
				case ">":
					currPosition = setCurrentPosition(currPosition, 0, 1);
					break;
				case "<":
					currPosition = setCurrentPosition(currPosition, 0, -1);
					break;
				default:
					System.out.println("DEFAULT");
					break;
			}

			if (searchPosition(currPosition, positions)) {
				continue;
			} else {
				positions.add(currPosition.clone());
			}
		}
		System.out.println(positions.size());
	}

	private static int[] setCurrentPosition(int[] curr, int up, int side) {
		curr[0] += up;
		curr[1] += side;
		return curr;
	}

	private static boolean searchPosition(int[] positionToFind, ArrayList<int[]> positions) {

		for (int i = 0; i < positions.size(); i++) {
			if (positionToFind[0] == positions.get(i)[0] && positionToFind[1] == positions.get(i)[1]) {
				return true;
			}
		}
		return false;
	}

	private static String stringFromFile(String filePath) throws IOException {
		try {
			BufferedReader reader = new BufferedReader(new FileReader(filePath));
			String line = reader.readLine();

			reader.close();
			return line;
		} catch (Exception e) {
			throw new IOException("Cannot find file");
		}
	}
}
