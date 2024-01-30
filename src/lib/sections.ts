export interface SectionType {
	desc: string;
	placeholder: string;
	useCase: string;
	type: string;
}

export const sections: SectionType[] = [
	{
		desc: 'Enter the url you would like to shorten',
		placeholder: 'URL...',
		useCase: 'url',
		type: 'text'
	},
	{
		desc: 'What do you want the path to be?',
		placeholder: 'Path...',
		useCase: 'path',
		type: 'text'
	},
	{
		desc: 'How long would you like the url to last?',
		placeholder: 'Time in days...',
		useCase: 'days',
		type: 'number'
	}
];
