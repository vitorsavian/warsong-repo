function Print_help()
	print("Uso: owlbear [comando] [opções]")
	print("Comandos:")
	print("  create character   Cria um novo personagem")
	print("Opções:")
	print("  --name <nome>      Define o nome do personagem")
	print("  --help             Exibe esta mensagem de ajuda")
end

function Main(args)
	if #args < 1 then
		Print_help()
		return
	end

	if args[1] == "create" and args[2] == "character" then
		local name = nil

		for i = 3, #args do
			if args[i] == "--name" then
				if i + 1 <= #args then
					name = args[i + 1]
				else
					print("Erro: --name requer um valor.")
					return
				end
			elseif args[i] == "--help" then
				Print_help()
				return
			else
				print("Erro: opção desconhecida " .. args[i])
				return
			end
		end

		if name then
			create_character(name) -- Chama a função C
		else
			print("Erro: Nome do personagem não fornecido.")
		end
	else
		Print_help()
	end
end

-- Chama a função principal passando os argumentos da linha de comando
Main(arg)
