export interface Language {
	name: string
	icon: {
		name: string,
		color?: string
	}
	supported: boolean
	concepts: number
	exercises: number
	url: string
	category?: "language" | "framework" | "tool"
}
