package day02;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.stream.Stream;

public class A {
	public static void main(String[] args) throws IOException {
		String filePath = "./input.txt";
		ArrayList<String> allLines = stringsFromLinesOfFile(filePath);

		int sum = 0;
		for (int i = 0; i < allLines.size(); i++) {
			ArrayList<Integer> metricsArr = new ArrayList<Integer>(Arrays.asList(
					Integer.parseInt(allLines.get(i).split("x")[0]),
					Integer.parseInt(allLines.get(i).split("x")[1]),
					Integer.parseInt(allLines.get(i).split("x")[2])));

			// length, width, height
			int withoutSlack = (2 * metricsArr.get(0) * metricsArr.get(1)) + (2 * metricsArr.get(1) * metricsArr.get(2))
					+ (2 * metricsArr.get(2) * metricsArr.get(0));

			metricsArr.sort(Comparator.naturalOrder());
			sum += withoutSlack + (metricsArr.get(0) * metricsArr.get(1));
		}
		System.out.println(sum);
	}

	private static ArrayList<String> stringsFromLinesOfFile(String filePath) throws IOException {
		try {
			BufferedReader reader = new BufferedReader(new FileReader(filePath));
			Stream<String> linesStream = reader.lines();

			ArrayList<String> allLines = new ArrayList<String>();
			linesStream.forEach(line -> allLines.add(line));

			linesStream.close();
			reader.close();
			return allLines;
		} catch (Exception e) {
			throw new IOException("cannot find file");
		}
	}
}
