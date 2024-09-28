function CharacterHelp()
	print("Usage: owlbear character [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create character")
	print("  update     Update character")
	print("  get        Get character")
	print("  delete     Delete character")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterCreateHelp()
	print("Usage: owlbear character create [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterUpdateHelp()
	print("Usage: owlbear character update [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterGetHelp()
	print("Usage: owlbear character get [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterDeleteHelp()
	print("Usage: owlbear character delete [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterHandler(args)
	for i = 2, #args do
		if args[i] == "--help" then
			CharacterHelp()
			return
		end
	end
end
