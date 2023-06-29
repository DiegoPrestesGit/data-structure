package day01;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class B {

	/**
	 *
	 * @param args
	 * @throws IOException
	 */
	public static void main(String[] args) throws IOException {
		String fileName = "./input.txt";
		String contents = stringFromFile(fileName);
		String[] strArr = contents.split("");

		int floorsCount = 0;
		int currentFloor = 0;

		for (int i = 0; i < strArr.length; i++) {
			floorsCount++;
			if (strArr[i].equals("(")) {
				currentFloor++;
			} else {
				currentFloor--;
			}
			if (currentFloor == -1) {
				System.out.println(floorsCount);
				break;
			}
		}
	}

	private static String stringFromFile(String filePath) throws IOException {
		try {
			BufferedReader reader = new BufferedReader(new FileReader(filePath));
			String line = reader.readLine();

			reader.close();
			return line;
		} catch (Exception e) {
			throw new IOException("cannot find file");
		}
	}
}
